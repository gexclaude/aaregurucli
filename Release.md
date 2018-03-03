# Releasing

Precondition - you need access to the github repository as well as to the corresponding homebrew-tap repository. You should have an access token and have it exported

    export GITHUB_TOKEN=...

If you have everything in place, start releasing

    ./release.sh -v v<x>.<y>.<z>
    
Whereas <x> is major, <y> minor and <z> bugfix version
