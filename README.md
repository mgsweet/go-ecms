# go-ecms
Golang error code management system. For error code doc, see: https://mgsweet.com/go-ecms/

## Preconditions
If you want to build site, install Hugo. For mac with homebrew:
```shell
brew install hugo
```
For other platforms, see: https://gohugo.io/getting-started/installing/

## Usage
To generate both go code and site code, run:
```shell
./build.sh
```

To build go code only, run:
```shell
./build_go.sh
```

To build site only, run:
```shell
./build_site.sh
```

To run hugo site in development mode:
```shell
cd ecms-site
hugo server
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