# go example

| Branch | Travis | Report.ci |
|--------|--------|-----------|
| develop | [![Build Status](https://travis-ci.com/report-ci/go-example.svg?branch=develop)](https://travis-ci.com/report-ci/go-example) | [![Build Status](https://api.report.ci/status/report-ci/go-example/badge.svg?branch=develop&level=cases)](https://api.report.ci/status/report-ci/go-example?branch=develop) |
| master  | [![Build Status](https://travis-ci.com/report-ci/go-example.svg?branch=master)](https://travis-ci.com/report-ci/go-example)  | [![Build Status](https://api.report.ci/status/report-ci/go-example/badge.svg?branch=master&level=cases)](https://api.report.ci/status/report-ci/go-example?branch=master) |

This repository demonstrates how to use report.ci with go. It can be done by adding one command to travis-ci as shown below.

```yml
after_script:
  - python <(curl -s https://report.ci/upload.py)
```

There's a example PR [here](https://github.com/report-ci/go-example/pull/1) and use badges in the markdown file like this:

```md
Last build (the last uploaded one, my be different than from the CI)

[![Build Status](https://api.report.ci/status/report-ci/go-example/badge.svg)](https://api.report.ci/status/report-ci/go-example)

Branch master

[![Build Status](https://api.report.ci/status/report-ci/go-example/badge.svg&branch=master)](https://api.report.ci/status/report-ci/go-example&branch=master)

You can also set the level to suites, the default is cases.

[![Build Status](https://api.report.ci/status/report-ci/go-example/badge.svg?branch=master&level=suites)](https://api.report.ci/status/report-ci/go-example?branch=master)

```

