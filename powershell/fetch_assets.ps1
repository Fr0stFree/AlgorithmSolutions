$token = # Your token
$userId = # Your user ID
$url = "https://mdr.kaspersky.com/api/v1/$userId/assets/list"
$headers = @{
    Authorization = "Bearer $token"
    ContentType = "application/json"
}
$save_path = "output.csv"
$max_page = 1000
$page_size = 10

for ($page = 1; $page -le $max_page; $page++) {
    $body = @{
        sort = "last_seen:asc"
        page_size = $page_size
        page = $page
    }

    $response = Invoke-WebRequest -Uri $url -Body ($body | ConvertTo-Json) -Headers $headers -Method POST
    if ($response.StatusCode -ne 200) {
        Write-Output "Failed to fetch assets. Exiting..."
        break
    }

    $body = $response.Content | ConvertFrom-Json
    $result = @()

    if ($body.Count -eq 0) {
        Write-Output "No more assets to fetch. Exiting..."
        break
    }
    foreach ($asset in $body) {
        $networkInfo = $asset.network_interfaces | ForEach-Object {
            $_.dsc,
            $_.dnsd,
            $_.defg,
            $_.mac,
            $_.ipcm,
            $_.ip
        }

        $result += [PSCustomObject]@{
            host_name               = $asset.host_name
            last_seen               = [datetime]::FromFileTimeUtc($asset.last_seen)
            first_seen              = [datetime]::FromFileTimeUtc($asset.first_seen)
            os_version              = $asset.os_version
            installed_product_info  = $asset.installed_product_info
            tenant_name             = $asset.tenant_name
            isolation               = $asset.isolation
            status                  = $asset.status
            ksc_host_id             = $asset.ksc_host_id
            asset_id                = $asset.asset_id
            domain                  = $asset.domain
            dsc                     = $networkInfo[0]
            dnsd                    = $networkInfo[1]
            defg                    = $networkInfo[2]
            mac                     = $networkInfo[3]
            ipcm                    = $networkInfo[4]
            ip                      = $networkInfo[5]
        }
    }
    $result | Export-Csv -Path $save_path -NoTypeInformation -Append -Encoding UTF8
    Write-Host "Page $page is fetched and saved to $save_path"
}
