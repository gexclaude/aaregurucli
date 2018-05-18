package outkonjunktiv

type substitution struct {
  original string
  replacement string
}

func getSubstitutions() []substitution {
  return []substitution{
    substitution{original: "bringt", replacement: "brächt"},
    substitution{original: "bringts", replacement: "brächts"},
    substitution{original: "cha", replacement: "chönnt"},
    substitution{original: "chani", replacement: "chönnti"},
    substitution{original: "findy", replacement: "fändi"},
    substitution{original: "ga", replacement: "gieng"},
    substitution{original: "gange", replacement: "gieng"},
    substitution{original: "isch", replacement: "wär"},
    substitution{original: "ist", replacement: "wär"},
    substitution{original: "macht", replacement: "miech"},
    substitution{original: "schwümme", replacement: "schwümmti"},
    substitution{original: "schwümmt", replacement: "schwümmti"},
    substitution{original: "tuet", replacement: "tüengi"},
    substitution{original: "wird", replacement: "würd"},
    substitution{original: "wot", replacement: "würd wöue"},
    substitution{original: "wott", replacement: "würd wöue"},
    substitution{original: "wüsse", replacement: "wüsste"},
  }
}