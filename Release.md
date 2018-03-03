# Releasing

Precondition - you need access to the github repository as well as to the corresponding homebrew-tap repository. You should have an access token and have it exported

    export GITHUB_TOKEN=...

If you have everything in place, start releasing

    git tag -a v<x>.<y>.<z>
    git push origin v<x>.<y>.<z>
    goreleaser