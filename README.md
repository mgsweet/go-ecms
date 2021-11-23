# go-ecms
Golang error code management system. For error code doc, see: https://mgsweet.com/go-ecms/

## Usage
To generate both go code and site code, run:
```shell
./build.sh
```

To run hugo site in development mode:
```shell
cd ecms-site
hugo server -D
```

Build:
```shell
cd ecms-site
hugo -d ../docs/
```

For Deployment, remember to change the `baseURL` in `/ecms-site` to your site:
```toml
baseURL = "https://your-site/"
```

## Credit
Hugo Template: https://geekdocs.de/