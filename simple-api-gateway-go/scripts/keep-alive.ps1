param(
    [string]$GatewayUrl = "http://localhost:8080/health",
    [int]$IntervalSeconds = 25,
    [int]$StartWaitSeconds = 5
)

$ErrorActionPreference = "Stop"

function Get-GatewayProcess {
    return Get-CimInstance Win32_Process |
        Where-Object { $_.Name -eq "go.exe" -and $_.CommandLine -match "./cmd/server" } |
        Select-Object -First 1
}

function Start-Gateway {
    Write-Host "[$(Get-Date -Format o)] Iniciando API Gateway..."
    $proc = Start-Process -FilePath "go" -ArgumentList "run ./cmd/server" -WorkingDirectory (Get-Location) -PassThru -WindowStyle Hidden
    Start-Sleep -Seconds $StartWaitSeconds
    return $proc
}

function Stop-Gateway {
    param([int]$ProcessId)

    if ($ProcessId -gt 0) {
        try {
            Stop-Process -Id $ProcessId -Force -ErrorAction Stop
            Write-Host "[$(Get-Date -Format o)] Proceso detenido (PID: ${ProcessId})."
        } catch {
            Write-Host "[$(Get-Date -Format o)] No se pudo detener PID ${ProcessId}: $($_.Exception.Message)"
        }
    }
}

Write-Host "[$(Get-Date -Format o)] Watchdog iniciado. Health endpoint: $GatewayUrl"

$running = Get-GatewayProcess
if (-not $running) {
    $started = Start-Gateway
    $currentPid = $started.Id
} else {
    $currentPid = $running.ProcessId
    Write-Host "[$(Get-Date -Format o)] API Gateway ya estaba activo (PID: $currentPid)."
}

while ($true) {
    try {
        $resp = Invoke-WebRequest -Uri $GatewayUrl -Method Get -TimeoutSec 8
        if ($resp.StatusCode -ne 200) {
            throw "Health check devolvio status $($resp.StatusCode)"
        }

        Write-Host "[$(Get-Date -Format o)] OK (200)"
    } catch {
        Write-Host "[$(Get-Date -Format o)] Health check fallo: $($_.Exception.Message)"

        $proc = Get-GatewayProcess
        if ($proc) {
            Stop-Gateway -ProcessId $proc.ProcessId
        } elseif ($currentPid) {
            Stop-Gateway -ProcessId $currentPid
        }

        $restarted = Start-Gateway
        $currentPid = $restarted.Id
    }

    Start-Sleep -Seconds $IntervalSeconds
}
