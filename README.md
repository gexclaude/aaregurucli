# aare.guru CLI

<a href="https://aare.guru/"><img src="https://github.com/gexclaude/aaregurucli/blob/master/docs/guru-logo-2018-3.png" alt="aare.guru" width="25%" /></a>

* [Was chani mit dr aare.guru CLI äpp?](#was-chani-mit-dr-aareguru-cli-%C3%A4pp)
* [Lifere statt lafere](#lifere-statt-lafere)
* [Grundsätzlechs](#grunds%C3%A4tzlechs)
* [Installation](#installation)
   * [Mac: Mit brew aareguru kömerle](#mac-mit-brew-aareguru-k%C3%B6merle)
   * [Linux](#linux)
      * [RPM: D RPM Datei häre chnüble](#rpm-d-rpm-datei-h%C3%A4re-chn%C3%BCble)
      * [Debian: D Deb Datei mit blose Chlööpe häre tue](#debian-d-deb-datei-mit-blose-chl%C3%B6%C3%B6pe-h%C3%A4re-tue)
   * [Windows](#windows)
      * [Windows: Mit scoop aareguru ufläse](#windows-mit-scoop-aareguru-ufl%C3%A4se)
      * [Windows: ds Päckli vo Häntsche abelade](#windows-ds-p%C3%A4ckli-vo-h%C3%A4ntsche-abelade)
* [Konfiguration](#konfiguration)
   * [Proxy](#proxy)
   * [Autocompletion](#autocompletion)
      * [Bash](#bash)
      * [Zsh](#zsh)
* [Merci's (Credits)](#mercis-credits)

## Aare.Guru?

Wär isch de dä aare.guru? Lue hie: https://aare.guru

## Was chani mit dr aare.guru CLI äpp?

Mit dere chasch ganz komod d Aare-Tämperatur, -Wassermängi u ds aktuelle bärner Wätter i dire Befählszyle abfragä - u das i gwaneter aare.guru Qualität.

Obenuse, nid?

## Liefere statt lafere

<img src="https://github.com/gexclaude/aaregurucli/blob/master/docs/gifs/aareguru-standard.gif" alt="aare.guru - schribmaschine" />

<img src="https://github.com/gexclaude/aaregurucli/blob/master/docs/gifs/aareguru-schribmaschine.gif" alt="aare.guru - schribmaschine" />

## Grundsätzlechs

Usfüüährig uf dr Kommandozile wie fougt:

    aareguru

Spicke duesch grundsätzläch so:

    aareguru --help

## Installation

### Mac: Mit brew aareguru kömerle

    brew tap gexclaude/homebrew-tap
    brew install aareguru

Aktualisierige duesch so

    brew upgrade aareguru

### Linux

Hie d'Variante für Pinguin-Fründe

### RPM: D RPM Datei häre chnüble

Suechdr die richtig Plattform us kopier dr Link zur RPM-Datei vodr Release Site [Releases](https://github.com/gexclaude/aaregurucli/releases)

    sudo rpm -Uvh https://github.com/gexclaude/aaregurucli/releases/download/v<...>/aareguru_<...>.rpm

### Debian: D Deb Datei mit blose Chlööpe häre tue

Suechdr die richtig Plattform us kopier dr Link zur Deb-Datei vodr Release Site [Releases](https://github.com/gexclaude/aaregurucli/releases)

    wget https://github.com/gexclaude/aaregurucli/releases/download/v<...>/aareguru_<...>.deb -0 aareguru.deb
    sudo dpkg -i aareguru.deb

### Windows

Kei Grund Befählszyle nid ds bruuche!

#### Windows: Mit scoop aareguru ufläse

Scoop muesch installiert ha:
http://scoop.sh/

Mit Powershell 3 geit das so:

    iex (new-object net.webclient).downloadstring('https://get.scoop.sh')
    
U faus de ne Proxy bruchsch machsches eso:

    $browser = New-Object System.Net.WebClient
    $browser.Proxy.Credentials =[System.Net.CredentialCache]::DefaultNetworkCredentials
    iex (new-object net.webclient).downloadstring('https://get.scoop.sh')
    scoop config proxy <host>:<port>

Je nachdem hiuft dr dr CNTLM Proxy: http://cntlm.sourceforge.net/
    
Denn duesch dr Chessu drzue und laschs la tschädere

    scoop bucket add aareguru https://github.com/gexclaude/scoop-bucket
    scoop install aareguru

Zum spöter aktualisiere

    scoop update aareguru

#### Windows: ds Päckli vo Häntsche abelade

Geisch zu de [Releases](https://github.com/gexclaude/aaregurucli/releases) und ladschdr ds Zip File für dini Plattform abe,
entpacksches nöime und laschs loufe. Viläch wosches ja o zum PATH hinzuefüege. Das muesch aber bi dere Methode säuber mache 

## Konfiguration

### Proxy

Wede e proxy bruchsch

    aareguru --proxy http://<host>:<port>

### Autocompletion

### Bash

Fougende Befähl spuckt s autcomplete script für Bash use:

    aareguru --completion-script-bash

Das chame entweder is `bash_profile` ine due (bevorzugt)

    eval "$(aareguru --completion-script-bash)"

oder fix hingerlege, de muesches aber säuber aktuell haute

    aareguru --completion-script-bash > aareguru_completion.sh
    sudo mv aareguru_completion.sh /etc/bash_completion.d/aareguru

### Zsh

Fougende Befähl spuckt s autcomplete script für Zsh use:

    aareguru --completion-script-zsh

## Merci's (Credits)

* https://github.com/goreleaser/goreleaser
* https://github.com/logrusorgru/aurora
* https://github.com/gosuri/uiprogress
* https://github.com/alecthomas/kingpin
* https://github.com/lukesampson/scoop
* https://github.com/Homebrew/brew/
