# aare.guru CLI

## Aare.Guru?

Wär isch de dä aare.guru? Lue hie: https://aare.guru

## Was chani mit dr aare.guru CLI äpp?

Mit dere chasch ganz komod d Aare-Tämperatur und Wassermängi i dire befälszyle abfragä - u das i gwaneter aare.guru qualität.

Obenuse, nid?

 
## Mac: Mit brew aareguru kömerle

    brew tap gexclaude/homebrew-tap
    brew install aareguru

u zwüschem `git commit` und em `git push` machsch hurti es `aareguru` :)

## Debian: D Deb Datei mit blose Chlööpe häre tue

Suechdr die richtig Plattform us kopier dr Link zur Deb-Datei vodr Release Site [Releases](https://github.com/gexclaude/aaregurucli/releases)

    wget https://github.com/gexclaude/aaregurucli/releases/download/v.../aareguru_<...>.deb -0 aareguru.deb
    sudo dpkg -i aareguru.deb

## Windows: Mit scoop aareguru ufläse

Scoop muesch installiert ha:
http://scoop.sh/

Mit Powershell 3 geit das so:

    iex (new-object net.webclient).downloadstring('https://get.scoop.sh')
    
Denn duesch dr Chessu drzue und laschs la tschädere

    scoop bucket add aareguru https://github.com/gexclaude/scoop-bucket
    scoop install aareguru

## Windows: ds Päckli vo Häntsche abelade

Geisch zu de [Releases](https://github.com/gexclaude/aaregurucli/releases) und ladschdr ds Zip File für dini Plattform abe,
entpacksches nöime und laschs loufe. Viläch wosches ja o zum PATH hinzuefüege. Das muesch aber bi dere Methode säuber mache 


## Merci's (Credits)

* https://github.com/goreleaser/goreleaser
* https://github.com/logrusorgru/aurora
* https://github.com/lukesampson/scoop
* https://github.com/Homebrew/brew/