# logrole

Logrole is a faster, usable, fine-grained client for exploring your Twilio
logs.

<img src="https://kev.inburke.com/rawblog/images/logrole.png" />

- Customizable permissions for each user browsing the site - limit access to
SMS/MMS bodies, resources older than a certain age, recordings, calls, call
from, etc. etc.

- Account Sid is hidden from end users at all times.

- Easy site search - tab complete and search for a sid to go straight to the
  instance view for that resource.

- Click-to-copy sids and phone numbers.

- MMS messages are always fetched over HTTPS. The default Twilio API/libraries
hand back insecure image links, but we rewrite URLs before fetching them.

- Tab to search: start typing the URL in the tab bar, then press &lt;tab&gt;.
  Paste any SID to immediately jump to that page.

## It Is Really Fast

Logrole fetches and caches the first page of every result set every 30 seconds,
and any time you page through records, the next page is prefetched and cached
before you click on it. This means paging through your Twilio logs via Logrole
is *significantly faster* than viewing results in your Dashboard or via the
API! If you don't believe me, the request latencies are displayed at the bottom
of every page; they're frequently 10ms and very rarely above 200ms.

Logrole uses hand-written HTML, one CSS file, one font, and close to no
Javascript, so render times and memory performance are very good.

If you need to search your Twilio Logs, this is the tool you should be using.

## Installation

To install Logrole, run

```bash
go get -u github.com/saintpete/logrole/...
```

You will need a working [Go environment][go-env]; I recommend setting the
GOPATH environment variable to `$HOME/go` in your .bashrc or equivalent.

```bash
export GOPATH="$HOME/go"
```

## Configuration and Deployment

There are two main ways to deploy Logrole. Either:

- Write all settings to a `config.yml` file (a sample is in
[config.sample.yml][config-sample]), then run `logrole_server
--config=config.yml`.

Or:

- Set all configuration as environment variables. Run
`logrole_write_config_from_env > config.yml` to write all those environment
variables to a config.yml file. Follow the steps in (1).

For more information, please [see the Settings documentation][settings-docs].

[settings-docs]: https://github.com/saintpete/logrole/blob/master/docs/settings.md
[config-sample]: https://github.com/saintpete/logrole/blob/master/config.sample.yml

## Authentication

Logrole supports three authentication modes: none, basic auth,
and Google OAuth. For more information, [see the Settings
documentation][settings-auth-docs].

[settings-auth-docs]: https://github.com/saintpete/logrole/blob/master/docs/settings.md#authentication

## Local Development

Logrole is written in Go; you'll need a [working Go environment][go-env]
running at least Go 1.7. Follow the instructions here to set one up:
https://golang.org/doc/install.

[go-env]: https://golang.org/doc/install

To check out the project, run `go get -u github.com/saintpete/logrole/...`.

To start a development server, run `make serve`. This will start a server on
[localhost:4114](http://localhost:4114).

By default we look for a config file in `config.yml`. The values for this
config file can be found in the FileConfig struct in `config/settings.go`.
There's an example config file at `config.sample.yml`.

## Run the tests

Please run the tests before submitting any changes. Run `make test` to run the
tests, or run `make race-test` to run the tests with the race detector enabled.

## View the documentation

Run `make docs`.

## Errata

The Twilio Dashboard displays Participants for completed Conferences, but [this
functionality is not available via the API][issue-4]. Please [contact Support
to request this feature][support] if you'd like it to be available in Logrole.

[support]: mailto:help@twilio.com
[issue-4]: https://github.com/saintpete/logrole/issues/4

The Start/End date filters may only work in Chrome.
