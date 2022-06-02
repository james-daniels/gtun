# tunnel

Tunnel is a simple wrapper for the gcloud iap tunneling command.

The purpose is to replace the following code:

```bash
gcloud compute start-iap-tunnel servername 22 --local-host-port=localhost:2222 --zone us-east-4
```

to the following:

```bash
tunnel up --linux
```

## Usage

### Tunnel Commands

There are only two options with tunnel, up and down

```text
$ tunnel -h
A simple app that uses the gcloud cli to establish iap tunnels.

Usage:
  tunnel [command]

Available Commands:
  down        Terminate the tunnel sessions
  help        Help about any command
  up          Establish a tunnel session

Flags:
  -h, --help   help for tunnel

Use "tunnel [command] --help" for more information about a command.
```

### Tunnel Up

Establishing a tunnel is as simple as using the following flags: -l, --linux, -w, --windows

If no flags are entered, both linux and windows tunnels will be established.

```text
$ tunnel up -h
Establish the linux or windows tunnels or all tunnels simultaneously

Usage:
  tunnel up [flags]

Flags:
  -h, --help      help for up
  -l, --linux     bring up linux tunnel
  -w, --windows   bring up windows tunnel
```

### Tunnel Down

Brining down a tunnel is a simple as using following flags: -l, --linux, -w, --windows, -a, --all

Unlike bringing up a tunnel, you must enter a flag to bring a tunnel down.

```text
$ tunnel down -h
Terminate the linux or windows tunnels or all tunnels simultaneously

Usage:
  tunnel down [flags]

Flags:
  -a, --all       bring down all tunnels
  -h, --help      help for down
  -l, --linux     bring down linux tunnel
  -w, --windows   bring down windows tunnel
```

### Installation

Determine best place to store the app on your filesystem. The example assumes your home directory.

```text
mkdir -p app/tunnel
cp tunnel app/tunnel

# Make sure the app is in your $PATH
sudo ln -s app/tunnel/tunnel /usr/local/bin/tunnel
```

## Configuration

The program reads a config file (config.ini) which needs to be placed same folder location as the executable.

Pay attention to the "command" parameter.  It is assumed that your gcloud command will be located in different location than mine.  Please edit accordingly.

```text
[gcloud]
command = "/usr/local/bin/gcloud"
zone  = "us-west2-b"

[local]
linport = 2022
winport = 3389

[servers]
linserver = "gcp-linux-server"
winserver = "gcp-windows-server"
```
