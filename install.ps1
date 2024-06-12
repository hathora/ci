$repoOwner = "hathora"
$repoName = "ci"
$binaryName = "hathora"

# Get the latest release tag from GitHub
$latestRelease = (Invoke-WebRequest "https://api.github.com/repos/$repoOwner/$repoName/releases/latest" | ConvertFrom-Json).tag_name

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
$downloadUrl = "https://github.com/$repoOwner/$repoName/releases/download/$latestRelease/$binaryName-windows-$arch"

# Set the install directory
$installDir = "$env:LOCALAPPDATA\Microsoft\WindowsApps"

# Download the binary
Write-Host "Downloading $binaryName for $arch..."
Invoke-WebRequest -Uri $downloadUrl -OutFile "$binaryName.exe"

# Install the binary
Write-Host "Installing $binaryName to $installDir..."
Move-Item -Path "$binaryName.exe" -Destination $installDir

Write-Host "Installation complete!"
