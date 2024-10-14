# Enable strict mode and exit on error
Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"
$ProgressPreference = "SilentlyContinue"

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

# Set the install directory
$installDir = "$env:LOCALAPPDATA\Microsoft\WindowsApps"

# Set the download URL for the Windows binary based on the architecture and version
if ($env:HATHORA_CLI_VERSION) {
    $downloadUrl = "https://github.com/$repoOwner/$repoName/releases/download/$env:HATHORA_CLI_VERSION/$binaryName-windows-$arch.exe"
    Write-Host "Downloading and installing $binaryName version $env:HATHORA_CLI_VERSION for $arch to $installDir..."
} else {
    $downloadUrl = "https://github.com/$repoOwner/$repoName/releases/latest/download/$binaryName-windows-$arch.exe"
    Write-Host "Downloading and installing latest version of $binaryName for $arch to $installDir..."
}

# Download and install the binary directly to the destination
Invoke-WebRequest -Uri $downloadUrl -OutFile "$installDir\$binaryName.exe"

Write-Host "Installation complete!"
