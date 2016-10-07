# Set up the prompt
autoload -Uz colors
colors

autoload -Uz promptinit
promptinit
PROMPT="%{$fg[blue]%}%* %{$fg_no_bold[yellow]%}%~ %{$reset_color%}%# "

setopt histignorealldups sharehistory

# Use emacs keybindings even if our EDITOR is set to vi
bindkey -e

# Keep 100000 lines of history within the shell and save it to ~/.zsh_history:
HISTSIZE=100000
SAVEHIST=10000000
HISTFILE=~/.zsh_history

# history modern completion system
autoload -Uz compinit
compinit

export LSCOLORS=exfxcxdxbxegedabagacad
export LS_COLORS='di=34:ln=35:so=32:pi=33:ex=31:bd=46;34:cd=43;34:su=41;30:sg=46;30:tw=42;30:ow=43;30'
export ZLS_COLORS=$LS_COLORS

zstyle ':completion:*' auto-description 'specify: %d'
zstyle ':completion:*' completer _expand _complete _correct _approximate
zstyle ':completion:*' format 'Completing %d'
zstyle ':completion:*' group-name ''
zstyle ':completion:*' menu select=2
eval "$(dircolors -b)"
zstyle ':completion:*:default' list-colors ${(s.:.)LS_COLORS}
zstyle ':completion:*' list-colors 'di=34' 'ln=35' 'so=32' 'ex=31' 'bd=46;34' 'cd=43;34'
zstyle ':completion:*' list-prompt %SAt %p: Hit TAB for more, or the character to insert%s
zstyle ':completion:*' matcher-list '' 'm:{a-z}={A-Z}' 'm:{a-zA-Z}={A-Za-z}' 'r:|[._-]=* r:|=* l:|=*'
zstyle ':completion:*' menu select=long
zstyle ':completion:*' select-prompt %SScrolling active: current selection at %p%s
zstyle ':completion:*' use-compctl false
zstyle ':completion:*' verbose true

zstyle ':completion:*:*:kill:*:processes' list-colors '=(#b) #([0-9]#)*=0=01;31'
zstyle ':completion:*:kill:*' command 'ps -u $USER -o pid,%cpu,tty,cputime,cmd'


# enable color support of ls and also add handy aliases
if [ -x /usr/bin/dircolors ]; then
    test -r ~/.dircolors && eval "$(dircolors -b ~/.dircolors)" || eval "$(dircolors -b)"
    alias ls='ls --color=auto'
    #alias dir='dir --color=auto'
    #alias vdir='vdir --color=auto'

    alias grep='grep --color=auto'
    alias fgrep='fgrep --color=auto'
    alias egrep='egrep --color=auto'
fi


## option

setopt no_beep
setopt auto_cd
setopt auto_pushd
setopt pushd_ignore_dups
setopt extended_glob
setopt hist_ignore_space
setopt equals
setopt notify # バックグランドジョブの状態変化を通知
setopt list_packed
setopt extended_history # ヒストリに実行時間も保存
setopt share_history


## alias
alias la="ls -a"
alias ll="ls -l"
alias lla="ls -la"

alias aget="apt-get"
alias ach="apt-cache"

alias ga="git add"
alias gs="git status"
alias gc="git commit"
alias gd="git diff"
alias gl="git log"
alias gb="git branch"

alias ..="cd .."
alias ...="cd ../.."
alias cl="clear"
alias cdd="cd ~/Desktop"
alias sou="source "

alias sudo="sudo -E"
alias cpan-list="find `perl -e 'print \"@INC\"'` -name \"*.pm\" -print"

alias -g L='| less'
alias -g H='| head'
alias -g T='| tail'
alias -g G='| grep'
alias -g X='| xclip'
alias xo='xclip -o'

#tmux 256 color
alias tmux='tmux -2'


