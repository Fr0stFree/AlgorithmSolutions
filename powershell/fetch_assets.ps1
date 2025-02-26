$accessToken = $null
$refreshToken = $null
$tokensFile = "tokens.json"
$userId = "" # your user id
$url = "https://mdr.kaspersky.com/api/v1/$userId/assets/list"
$refreshTokenUrl = "https://mdr.kaspersky.com/api/v1/$userId/session/confirm"
$savePath = "output.csv"
$maxPage = 10000
$pageSize = 100


function LoadTokensFromFile {
    Write-Host "Loading tokens from file..."
    $tokens = Get-Content -Path $tokensFile | ConvertFrom-Json
    if ($null -eq $tokens.refresh_token) {
        Write-Host "Refresh token not found. Check that file '$tokensFile' contains 'refresh_token' field"
        Exit
    }
    Set-Variable -Name accessToken -Value $tokens.access_token -Scope Global
    Set-Variable -Name refreshToken -Value $tokens.refresh_token -Scope Global
    Write-Host "Tokens loaded"
}

function RefreshAccessToken {
    Write-Host "Refreshing access token..."
    $body = @{
        refresh_token = $refreshToken
    }
    try {
        $response = Invoke-WebRequest -Uri $refreshTokenUrl -Body ($body | ConvertTo-Json) -Method POST
    } catch {
        Write-Host "Failed to refresh access token. Response from server: $($_.Exception.Response)"
        Exit
    }
    $body = $response.Content | ConvertFrom-Json

    Set-Variable -Name accessToken -Value $body.access_token -Scope Global
    Set-Variable -Name refreshToken -Value $body.refresh_token -Scope Global
    $tokens = @{
        access_token = $accessToken
        refresh_token = $refreshToken
    }
    $tokens | ConvertTo-Json | Set-Content -Path $tokensFile
    Write-Host "Access token refreshed"
}

function FetchAssets {
    param (
        [int]$page
    )
     $headers = @{
         Authorization = "Bearer $accessToken"
         ContentType = "application/json"
    }
    $body = @{
        sort = "computer_name_hostname:asc"
        page_size = $pageSize
        page = $page
        version = 2
    }
    try {
        $response = Invoke-WebRequest -Uri $url -Body ($body | ConvertTo-Json) -Headers $headers -Method POST
    } catch {
        $statusCode = $_.Exception.Response.StatusCode.Value__
        if ($statusCode -eq 403) {
            RefreshAccessToken
            $headers.Authorization = "Bearer $accessToken"
            $response = Invoke-WebRequest -Uri $url -Body ($body | ConvertTo-Json) -Headers $headers -Method POST
        } else {
            Write-Host "Failed to retrieve assets. Response from server: $($_.Exception.Response)"
            Exit
        }
    }
    $body = $response.Content | ConvertFrom-Json

    if ($null -eq $body -or $body.Count -eq 0) {
        Write-Host "No assets found on page $page"
        return $null
    }

    $assets = @()
    foreach ($asset in $body) {
        $assetObj = [Asset]::new($asset)
        $assets += $assetObj
    }
    Write-Host "Page $page fetched and saved to $savePath"
    return $assets
}

function SaveAssets {
    param (
        [array]$assets
    )
    $assets | Export-Csv -Path $savePath -NoTypeInformation -Append -Encoding UTF8
}

function Main() {
    LoadTokensFromFile
    if ($null -eq $accessToken) {
        RefreshAccessToken
    }
    Write-Host "Fetching assets..."
    $null | Set-Content -Path $savePath 
    for ($page = 1; $page -le $maxPage; $page++) {
        $result = FetchAssets -page $page
        if ($null -eq $result) {
            break
        }
        SaveAssets -assets $result
    }
    Write-Host "Finished fetching assets"
}

Main

class Asset {
    [string] $host_name
    [datetime] $last_seen
    [datetime] $first_seen
    [string] $os_version
    [string] $installed_product_info
    [string] $tenant_name
    [string] $isolation
    [string] $status
    [string] $ksc_host_id
    [string] $asset_id
    [string] $domain
    [string] $dsc
    [string] $dnsd
    [string] $defg
    [string] $mac
    [string] $ipcm
    [string] $ip

    Asset([PSObject]$assetData) {
        $this.host_name = if ($null -eq $assetData.host_name) { "" } else { $assetData.host_name }
        $this.last_seen = [System.DateTimeOffset]::FromUnixTimeMilliseconds($assetData.last_seen).DateTime
        $this.first_seen = [System.DateTimeOffset]::FromUnixTimeMilliseconds($assetData.first_seen).DateTime
        $this.os_version = if ($null -eq $assetData.os_version) { "" } else { $assetData.os_version }
        $this.installed_product_info = if ($null -eq $assetData.installed_product_info) { "" } else { $assetData.installed_product_info }
        $this.tenant_name = if ($null -eq $assetData.tenant_name) { "" } else { $assetData.tenant_name }
        $this.isolation = if ($null -eq $assetData.isolation) { "" } else { $assetData.isolation }
        $this.status = if ($null -eq $assetData.status) { "" } else { $assetData.status }
        $this.ksc_host_id = if ($null -eq $assetData.ksc_host_id) { "" } else { $assetData.ksc_host_id }
        $this.asset_id = if ($null -eq $assetData.asset_id) { "" } else { $assetData.asset_id }
        $this.domain = if ($null -eq $assetData.domain) { "" } else { $assetData.domain }
        $this.dsc = ($assetData.network_interfaces.dsc | Where-Object { $_ -ne $null }) -join "|"
        $this.dnsd = ($assetData.network_interfaces.dnsd | Where-Object { $_ -ne $null }) -join "|"
        $this.defg = ($assetData.network_interfaces.defg | Where-Object { $_ -ne $null }) -join "|"
        $this.mac = ($assetData.network_interfaces.mac | Where-Object { $_ -ne $null }) -join "|"
        $this.ipcm = ($assetData.network_interfaces.ipcm | Where-Object { $_ -ne $null }) -join "|"
        $this.ip = ($assetData.network_interfaces.ip | Where-Object { $_ -ne $null }) -join "|"
    }
}