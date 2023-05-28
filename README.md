[![Go Report Card](https://goreportcard.com/badge/github.com/aceberg/git-confed)](https://goreportcard.com/report/github.com/aceberg/git-confed)
[![Binary-release](https://github.com/aceberg/git-confed/actions/workflows/release.yml/badge.svg)](https://github.com/aceberg/git-confed/actions/workflows/release.yml)
![GitHub](https://img.shields.io/github/license/aceberg/git-confed)
[![Maintainability](https://api.codeclimate.com/v1/badges/b326d121b6eb53713396/maintainability)](https://codeclimate.com/github/aceberg/git-confed/maintainability)

<h1><a href="https://github.com/aceberg/git-confed">
    <img src="https://raw.githubusercontent.com/aceberg/git-confed/main/assets/logo.png" width="20" />
</a>git-confed</h1>
<br/>

Overview all your git repos and edit `.git/config` with predefined blocks. For example, if you need to add `user` to some repos, just configure block with username and email and then click `Add block` in each repo.   
The app is capable of detecting git repos in folder.

- [Installation](https://github.com/aceberg/git-confed#installation)   
- [Usage](https://github.com/aceberg/git-confed#usage)   
- [Config](https://github.com/aceberg/git-confed#config)   
- [Options](https://github.com/aceberg/git-confed#options)  
- [Thanks](https://github.com/aceberg/git-confed#thanks)

![screenshot](https://raw.githubusercontent.com/aceberg/git-confed/main/assets/Screenshot%202023-05-26%20at%2016-08-42%20Git%20Config%20Editor.png)     

![screenshot](https://raw.githubusercontent.com/aceberg/git-confed/main/assets/Screenshot%202023-05-25%20at%2017-17-23%20Git%20Config%20Editor.png)

## Installation

### 1. From .deb repository (recommended)
```sh
curl -s --compressed "https://aceberg.github.io/ppa/KEY.gpg" | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/aceberg.gpg
```
```sh
sudo curl -s --compressed -o /etc/apt/sources.list.d/aceberg.list "https://aceberg.github.io/ppa/aceberg.list"
```
```sh
sudo apt update && sudo apt install git-confed
```
### 2. From .deb file
Download [latest](https://github.com/aceberg/git-confed/releases/latest) release, install with your package maneger

### 3. From .tar.gz
Download [latest](https://github.com/aceberg/git-confed/releases/latest) release, then
```sh
tar xvzf git-confed-*.tar.gz
cd git-confed
sudo ./install.sh
```

## Usage
### 1. Systemd as user (recommended)
Enable and start service, replace `MYUSER` with your username
```sh
sudo systemctl enable git-confed@MYUSER.service
sudo systemctl start git-confed@MYUSER.service
```
Web GUI will be available at [http://0.0.0.0:8848](http://0.0.0.0:8848)

### 2. Systemd as root
Enable and start service
```sh
sudo systemctl enable git-confed.service
sudo systemctl start git-confed.service
```
Web GUI will be available at [http://0.0.0.0:8848](http://0.0.0.0:8848)

### 3. From command line
Just run `git-confed`. Be mindful of the config files paths listed in [options](https://github.com/aceberg/git-confed#options) section.


## Config
### 1. With web GUI (recommended)
You can do all configuration through web interface. Config files paths are listed in [options](https://github.com/aceberg/git-confed#options) section below.

### 2. CLI
`blocks.yaml`
```yaml
Test Remote: "[remote \"test\"]\r\n\turl = bf:/home/data/repo/bare/testrepo"
User - aceberg: "[user]\r\n\tname = aceberg\r\n\temail = email@example.com"
```
`config.yaml`
```yaml
folders:
    - /home/data/repo/01-cloned
    - /home/data/repo/03-code
    - /home/data/repo/awesome
host: 0.0.0.0
other:
    - test
port: "8848"
theme: darkly
urls:
    - bitbucket
    - github
    - gitlab
```
| Variable  | Description | Default |
| --------  | ----------- | ------- |
| host | Address for web GUI | 0.0.0.0 |
| port | Port for web GUI | 8848 |
| theme | Any theme name from https://bootswatch.com in lowcase | darkly |
| folders | Where to search for git repos | |
| urls | Search for those in config file | bitbucket github gitlab |
| other | Search for those in config file | |



## Options

| Key  | Description | Default | Systemd (user) | Systemd (root) |
| --------  | ----------- | ------- | --- | --- |
| -b | Path to file with code blocks | /data/git-confed/blocks.yaml | $HOME/.config/git-confed/blocks.yaml | /etc/git-confed/blocks.yaml |
| -c | Path to GUI config file | /data/git-confed/config.yaml | $HOME/.config/git-confed/config.yaml | /etc/git-confed/config.yaml |


## Thanks
- All go packages listed in [dependencies](https://github.com/aceberg/git-confed/network/dependencies)
- [Bootstrap](https://getbootstrap.com/)
- Themes: [Free themes for Bootstrap](https://bootswatch.com)
- Icon <a href="https://www.flaticon.com/free-icons/gear" title="gear icons">Gear icons created by Anggara - Flaticon</a>