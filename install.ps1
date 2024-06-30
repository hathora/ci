# Enable strict mode and exit on error
Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"

$repoOwner = "hathora"
$repoName = "ci"
$binaryName = "hathora"

# Detect the architecture
$arch = $env:PROCESSOR_ARCHITECTURE
if ($arch -eq "AMD64") {
    $arch = "amd64"
} elseif ($arch -eq "ARM64") {
    $arch = "arm64"
} else {
    Write-Error "Unsupported architecture: $arch"
    exit 1
}

# Set the download URL for the Windows binary based on the architecture
$downloadUrl = "https://github.com/$repoOwner/$repoName/releases/latest/download/$binaryName-windows-$arch.exe"

# Set the install directory
$installDir = "$env:LOCALAPPDATA\Microsoft\WindowsApps"

# Download and install the binary directly to the destination
Write-Host "Downloading and installing $binaryName for $arch to $installDir..."
Invoke-WebRequest -Uri $downloadUrl -OutFile "$installDir\$binaryName.exe"

Write-Host "Installation complete!"
