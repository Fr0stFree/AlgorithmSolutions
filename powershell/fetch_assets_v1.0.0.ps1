param (
    [string]$userId = "",
    [string]$tokensFile = "tokens.json",
    [string]$savePath = "output.csv",
    [int]$startPage = 1,
    [int]$pageSize = 5000,
    [string]$baseUrl = "https://mdr.kaspersky.com",
    [int]$maxRetries = 5
)

$accessToken = $null
$refreshToken = $null
$url = $baseUrl + "/api/v1/$userId/assets/list"
$refreshTokenUrl = $baseUrl + "/api/v1/$userId/session/confirm"

function WriteLog {
    param (
        [string]$message,
        [string]$level = "INFO"
    )
    switch ($level) {
        "INFO" { $color = "White" }
        "WARNING" { $color = "Yellow" }
        "ERROR" { $color = "Red" }
    }
    $timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
    Write-Host "$timestamp [$level] $message" -ForegroundColor $color
}

# Загружает токены доступа из файла. Проверяет наличие refresh token.
function LoadTokensFromFile {
    WriteLog "Loading tokens from file..."
    $tokens = Get-Content -Path $tokensFile | ConvertFrom-Json
    if ($null -eq $tokens.refresh_token) {
        WriteLog "Refresh token not found. Check that file '$tokensFile' contains 'refresh_token' field" -level "ERROR"
        Exit 1
    }
    Set-Variable -Name accessToken -Value $tokens.access_token -Scope Global
    Set-Variable -Name refreshToken -Value $tokens.refresh_token -Scope Global
    WriteLog "Tokens loaded"
}

# Обновляет access token с использованием refresh token. Сохраняет новые токены обратно в файл.
function RefreshAccessToken {
    WriteLog "Refreshing access token..."
    $body = @{
        refresh_token = $refreshToken
    }
    try {
        $response = Invoke-WebRequest -Uri $refreshTokenUrl -Body ($body | ConvertTo-Json) -Method POST
    }
    catch {
        WriteLog "Failed to refresh access token, an error occurred: $($_.Exception.Message)" -level "ERROR"
        Exit 1
    }
    $body = $response.Content | ConvertFrom-Json

    Set-Variable -Name accessToken -Value $body.access_token -Scope Global
    Set-Variable -Name refreshToken -Value $body.refresh_token -Scope Global
    $tokens = @{
        access_token  = $accessToken
        refresh_token = $refreshToken
    }
    $tokens | ConvertTo-Json | Set-Content -Path $tokensFile
    WriteLog "Access token refreshed"
}

# Получает данные об активах с указанной страницы. Принимает параметр:
# $page - номер страницы для запроса
# Автоматически обновляет access token, если получен код ответа 403 (Forbidden).
function FetchAssets {
    param (
        [int]$page,
        [int]$currentTry = 0
    )
    $headers = @{
        Authorization = "Bearer $accessToken"
        ContentType   = "application/json"
    }
    $body = @{
        sort      = "first_seen:asc"
        page_size = $pageSize
        page      = $page
        version   = 2
    }
    $startTime = Get-Date
    try {
        $response = Invoke-WebRequest -Uri $url -Body ($body | ConvertTo-Json) -Headers $headers -Method POST
    }
    catch {
        WriteLog "Failed to retrieve assets, an error occurred: $($_.Exception.Message)" -level "WARNING"
        if ($currentTry -ge $maxRetries) {
            WriteLog "Maximum retry count has been reached" -level "ERROR"
            Exit 1
        }
        WriteLog "Retrying $($currentTry + 1) out of $maxRetries..."
        $statusCode = $_.Exception.Response.StatusCode.Value__
        if ($statusCode -eq 403) {
            RefreshAccessToken
        }
        return FetchAssets -page $page -currentTry ($currentTry + 1)
    }
    $body = $response.Content | ConvertFrom-Json

    if ($null -eq $body -or $body.Count -eq 0) {
        WriteLog "No assets found on page $page"
        return $null
    }

    $assets = @()
    foreach ($asset in $body) {
        $assetObj = [Asset]::new($asset)
        $assets += $assetObj
    }
    $endTime = Get-Date
    $duration = $endTime - $startTime
    WriteLog "Page $page fetched with $($assets.Count) assets in $($duration.TotalSeconds) seconds"
    return $assets
}

# Сохраняет полученные данные об активах в CSV-файл. Принимает параметр:
# $assets - массив объектов Asset для сохранения
function SaveAssets {
    param (
        [array]$assets
    )
    $assets | Export-Csv -Path $savePath -NoTypeInformation -Append -Encoding UTF8
}

# Основная функция, которая:
# Загружает токены. При необходимости обновляет access token. Получает данные об активах постранично. Сохраняет результаты в CSV
function Main() {
    if ("" -eq $userId) {
        WriteLog "User id must be specified using -userId parameter" -level ERROR
        Exit 1
    }
    LoadTokensFromFile
    if ($null -eq $accessToken) {
        RefreshAccessToken
    }
    WriteLog "Starting to fetching assets from page $startPage with size $pageSize..."
    $startTimeTotal = Get-Date
    $null | Set-Content -Path $savePath 
    for ($page = $startPage; $page -le 10000; $page++) {
        $result = FetchAssets -page $page
        if ($null -eq $result) {
            break
        }
        SaveAssets -assets $result
    }
    $endTimeTotal = Get-Date
    $durationTotal = $endTimeTotal - $startTimeTotal
    WriteLog "Finished to fetch assets. Elapsed time: $($durationTotal.TotalSeconds) seconds"
}

# Entrypoint
Main

class Asset {
    [string] $host_name                  # имя хоста
    [datetime] $last_seen                # дата последнего пинга хоста
    [datetime] $first_seen               # дата первого пинга хоста
    [string] $os_version                 # версия ОС
    [string] $installed_product_info     # информация об установленных продуктах
    [string] $tenant_name                # принадлежность хоста к тенанту
    [string] $isolation                  # заизолирован ли хост
    [string] $status                     # статус хоста
    [string] $ksc_host_id                # идентификатор хоста в KSC
    [string] $asset_id                   # уникальный идентификатор хоста
    [string] $domain                     # домен
    [string] $dsc                        # различные сетевые параметры (объединяются через "|" при наличии нескольких значений)
    [string] $dnsd
    [string] $defg
    [string] $mac
    [string] $ipcm
    [string] $ip

    Asset([PSObject]$assetData) {
        $this.host_name = if ($null -eq $assetData.host_name) { "" } else { $assetData.host_name.Trim() }
        $this.last_seen = [System.DateTimeOffset]::FromUnixTimeMilliseconds($assetData.last_seen).DateTime
        $this.first_seen = [System.DateTimeOffset]::FromUnixTimeMilliseconds($assetData.first_seen).DateTime
        $this.os_version = if ($null -eq $assetData.os_version) { "" } else { $assetData.os_version.Trim() }
        $this.installed_product_info = if ($null -eq $assetData.installed_product_info) { "" } else { $assetData.installed_product_info.Trim() }
        $this.tenant_name = if ($null -eq $assetData.tenant_name) { "" } else { $assetData.tenant_name.Trim() }
        $this.isolation = if ($null -eq $assetData.isolation) { "" } else { $assetData.isolation }
        $this.status = if ($null -eq $assetData.status) { "" } else { $assetData.status.Trim() }
        $this.ksc_host_id = if ($null -eq $assetData.ksc_host_id) { "" } else { $assetData.ksc_host_id.Trim() }
        $this.asset_id = if ($null -eq $assetData.asset_id) { "" } else { $assetData.asset_id.Trim() }
        $this.domain = if ($null -eq $assetData.domain) { "" } else { $assetData.domain.Trim() }
        $this.dsc = ($assetData.network_interfaces.dsc | Where-Object { $_ -ne $null }) -join "|"
        $this.dnsd = ($assetData.network_interfaces.dnsd | Where-Object { $_ -ne $null }) -join "|"
        $this.defg = ($assetData.network_interfaces.defg | Where-Object { $_ -ne $null }) -join "|"
        $this.mac = ($assetData.network_interfaces.mac | Where-Object { $_ -ne $null }) -join "|"
        $this.ipcm = ($assetData.network_interfaces.ipcm | Where-Object { $_ -ne $null }) -join "|"
        $this.ip = ($assetData.network_interfaces.ip | Where-Object { $_ -ne $null }) -join "|"
    }
}
