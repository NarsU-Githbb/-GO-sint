#!/bin/sh
set -e

IMAGE_NAME="digital-detective"

# Build Docker image only once
docker build -t "$IMAGE_NAME" .

mkdir -p reports

# Run all tests in one container
if [ "$1" = "-all" ]; then
    echo "Running all Digital Detective tests..."

    docker run --rm -it \
        --env-file .env \
        -v "$(pwd)/reports:/app/reports" \
        --entrypoint /bin/sh \
        "$IMAGE_NAME" -c '
            set +e

            echo "\n========== FULL NAME SEARCH =========="
            data-digger -n "john doe"

            echo "\n========== RAW FULL NAME SEARCH =========="
            data-digger -raw -n "john doe"

            echo "\n========== IP SEARCH =========="
            data-digger -ip 8.8.8.1

            echo "\n========== RAW IP SEARCH =========="
            data-digger -raw -ip 8.8.8.1

            echo "\n========== USERNAME SEARCH =========="
            data-digger -un torvalds

            echo "\n========== ALL TESTS COMPLETED =========="
        '

else
    docker run --rm -it \
        --env-file .env \
        -v "$(pwd)/reports:/app/reports" \
        "$IMAGE_NAME" "$@"
fi