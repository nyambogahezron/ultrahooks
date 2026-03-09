const fs = require('fs');
const https = require('https');
const path = require('path');
const os = require('os');
const { execSync } = require('child_process');

const VERSION = require('./package.json').version;
const REPO = 'nyambogahezron/ultrahooks';

function getReleaseAsset() {
    const platform = os.platform();
    const arch = os.arch();

    let goos = platform;
    let goarch = arch;

    // Map platforms
    if (platform === 'win32') goos = 'windows';

    // Map architectures
    if (arch === 'x64') goarch = 'amd64';
    else if (arch === 'arm64') goarch = 'arm64';

    let ext = '';
    if (goos === 'windows') {
        ext = '.exe';
    }

    // Format matches: ultrahooks_{os}_{arch}
    const assetName = `ultrahooks_${goos}_${goarch}${ext}`;
    return assetName;
}

function downloadBinary() {
    const assetName = getReleaseAsset();
    const binName = os.platform() === 'win32' ? 'ultrahooks.exe' : 'ultrahooks';
    const binPath = path.join(__dirname, 'bin', binName);
    
    // Ensure bin directory exists
    const binDir = path.dirname(binPath);
    if (!fs.existsSync(binDir)) {
        fs.mkdirSync(binDir, { recursive: true });
    }

    const url = `https://github.com/${REPO}/releases/download/v${VERSION}/${assetName}`;

    console.log(`Downloading ultrahooks v${VERSION} for ${os.platform()}-${os.arch()}...`);
    console.log(`URL: ${url}`);

    const file = fs.createWriteStream(binPath);

    function fetch(url) {
        https.get(url, (response) => {
            if (response.statusCode === 301 || response.statusCode === 302) {
                // Follow redirects
                return fetch(response.headers.location);
            }

            if (response.statusCode !== 200) {
                console.error(`Failed to download: ${response.statusCode} ${response.statusMessage}`);
                fs.unlinkSync(binPath);
                process.exit(1);
            }

            response.pipe(file);

            file.on('finish', () => {
                file.close();
                console.log('Download complete.');
                
                // Make binary executable on non-windows
                if (os.platform() !== 'win32') {
                    fs.chmodSync(binPath, 0o755);
                }
            });
        }).on('error', (err) => {
            fs.unlinkSync(binPath);
            console.error(`Error downloading binary: ${err.message}`);
            process.exit(1);
        });
    }

    fetch(url);
}

downloadBinary();
