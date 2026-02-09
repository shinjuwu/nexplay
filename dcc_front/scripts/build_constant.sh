# build constant

SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]:-$0}"; )" &> /dev/null && pwd 2> /dev/null; )";

cd "$SCRIPT_DIR"

cd ../buildJsFile

go run . ../frontend/src/base/common/constant.js ../frontend/src/base/common/featureCodes.js

cd ../frontend

npx prettier --write ./src/base/common/{constant,featureCodes}.js
