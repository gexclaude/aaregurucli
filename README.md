# aare.guru CLI

<a href="https://aare.guru/"><img src="https://github.com/gexclaude/aaregurucli/blob/master/docs/guru-logo-2018-3.png" alt="aare.guru" style="width: 50%" /></a>

## Aare.Guru?

Wär isch de dä aare.guru? Lue hie: https://aare.guru

## Was chani mit dr aare.guru CLI äpp?

Mit dere chasch ganz komod d Aare-Tämperatur, -Wassermängi u ds aktuelle bärner Wätter i dire Befählszyle abfragä - u das i gwaneter aare.guru Qualität.

Obenuse, nid?

## Grundsätzlechs

Usfüüährig uf dr Kommandozile wie fougt:

    aareguru

Spicke duesch grundsätzläch so:

    aareguru --help

Wede e proxy bruchsch

    aareguru --proxy http://<host>:<port>

## Mac: Mit brew aareguru kömerle

    brew tap gexclaude/homebrew-tap
    brew install aareguru

Aktualisierige duesch so

    brew upgrade aareguru

## RPM: D RPM Datei häre chnüble

Suechdr die richtig Plattform us kopier dr Link zur RPM-Datei vodr Release Site [Releases](https://github.com/gexclaude/aaregurucli/releases)

    sudo rpm -Uvh https://github.com/gexclaude/aaregurucli/releases/download/v<...>/aareguru_<...>.rpm

## Debian: D Deb Datei mit blose Chlööpe häre tue

Suechdr die richtig Plattform us kopier dr Link zur Deb-Datei vodr Release Site [Releases](https://github.com/gexclaude/aaregurucli/releases)

    wget https://github.com/gexclaude/aaregurucli/releases/download/v<...>/aareguru_<...>.deb -0 aareguru.deb
    sudo dpkg -i aareguru.deb

## Windows: Mit scoop aareguru ufläse

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

## Windows: ds Päckli vo Häntsche abelade

Geisch zu de [Releases](https://github.com/gexclaude/aaregurucli/releases) und ladschdr ds Zip File für dini Plattform abe,
entpacksches nöime und laschs loufe. Viläch wosches ja o zum PATH hinzuefüege. Das muesch aber bi dere Methode säuber mache 

## Merci's (Credits)

* https://github.com/goreleaser/goreleaser
* https://github.com/logrusorgru/aurora
* https://github.com/gosuri/uiprogress
* https://github.com/alecthomas/kingpin
* https://github.com/lukesampson/scoop
* https://github.com/Homebrew/brew/
