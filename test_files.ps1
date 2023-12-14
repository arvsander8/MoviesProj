# Define the root directory of your Go project
$rootDirectory = "F:\2023\Milkneko\REACT-GO\UDEMY\GraphQL\test-back-end"

# Define a recursive function to create test files
function CreateTestFiles {
    param (
        [string]$directory
    )

    # Get all Go source files in the current directory
    $goFiles = Get-ChildItem -Path $directory -Filter *.go -Recurse

    foreach ($file in $goFiles) {
        # Skip test files
        if ($file.Name -like "*_test.go") {
            continue
        }

        # Determine the test file name
        $testFileName = $file.DirectoryName + '\' + $file.BaseName + "_test.go"

        # Create the test file if it doesn't exist
        if (-not (Test-Path -Path $testFileName)) {
            New-Item -Path $testFileName -ItemType "file"
            "package " + $file.BaseName | Out-File -FilePath $testFileName
            Write-Host "Created test file: $testFileName"
        }
    }
}

# Start the recursive creation of test files
CreateTestFiles -directory $rootDirectory
