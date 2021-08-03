# hail
[![Go Report Card](https://goreportcard.com/badge/github.com/frzam/hail)](https://goreportcard.com/report/github.com/frzam/hail)
> Cross-Platform script management CLI written in go.

* [About](#about)
* [Features](#features)
* [Installation](#installation)
    * [Linux Or Unix](#linux-or-unix)
    * [Windows](#windows)
*  [Usage](#usage)
*  [License](#license)

## About
hail is a cross-platfrom script management CLI written in go. 

## Features
* Cross-Platform 
* Auto Completion scripts are available (Bash, fish, zsh and powershell)
* Inbuilt fuzzy searching
* Store commands as well as scripts
* Single and small binary to install. (Thanks to go!)
* Run commands or scripts directly from hail
* Portable toml file *.hailconfig*
* Use it as command line bookmark
* Format of scripts or command is also preserved eg *yaml* 

## Installation
### Linux or Unix
* Download the latest version of hail from releases https://github.com/frzam/hail/releases/<br>
    ```> wget https://github.com/frzam/hail/releases/download/v0.1.11/hail_0.1.11_Linux_x86_64.tar.gz```
* Unzip the tar.gz file <br>
    ```> tar -xf hail_0.1.11_Linux_x86_64.tar.gz```    
* Give execute permission to *hail*<br>
    ```> chmod +x hail```    
* Move the binary into *bin* folder so the it is accessible everywhere.<br>
    ```> mv hail /usr/local/bin/```
* Test if hail is working properly.<br>
    ```> hail version```
* Initialize hailconfig, it will create *.hailconfig* file under **$HOME**. If you want to create *.hailconfig* anywhere else then set env **HAILCONFIG** to that path.<br>
    ```> hail init <title>```
* [OPTIONAL] Set up tab auto completion for bash. Auto completion scripts are also available for fish and zsh.
    * Generate bash script in a file<br>
      ```> hail completion bash > ~/.hail```
    * Open ~/.bashrc in editor<br>
        ```> vi ~/.bashrc```
    * Add below line in .bashrc<br>
       ```> source ~/.hail```
     * Refresh .bashrc by doing<br>
        ```> source ~/.bashrc```  

### Windows
* Download the latest binary from releases https://github.com/frzam/hail/releases/ 
* Unzip the **hail__Windows_x86_64.tar.gz** file into *hail.exe*
* Place *hail.exe* into the *PATH*
* Test if hail is working properly. Open Command Prompt, Powershell or Git bash and run.<br>
    ```> hail version```
* Initialize hailconfig, it will create .hailconfig file under **$USERPROFILE**. If you want to create .hailconfig anywhere else then set env **HAILCONFIG** to that path<br>    ```> hail init <title>```
* Generate powershell completion script.<br>
```> hail completion powershell```