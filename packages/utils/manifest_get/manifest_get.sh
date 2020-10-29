#!/usr/bin/env bash
#Based on github.com/jasperes/bash-yaml

if [ "$#" -ne 1 ]; then
	>&2 echo "This script requires exactly 1 argument"
	exit 1
fi

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
MANIFEST=$SCRIPT_DIR/../../Manifest.yml
KEY=$1

SPACE='[[:space:]]*'
WORD='[a-zA-Z0-9_.-]*'
FS="$(echo @|tr @ '\034')"

RESULT=$(
	cat $MANIFEST |

	sed -e '/- [^\“]'"[^\']"'.*: /s|\([ ]*\)- \([[:space:]]*\)|\1-\'$'\n''  \1\2|g' |

	sed -ne '/^--/s|--||g; s|\"|\\\"|g; s/[[:space:]]*$//g;' \
		-e "/#.*[\"\']/!s| #.*||g; /^#/s|#.*||g;" \
		-e "s|^\($SPACE\)\($WORD\)$SPACE:$SPACE\"\(.*\)\"$SPACE\$|\1$FS\2$FS\3|p" \
		-e "s|^\($SPACE\)\($WORD\)${s}[:-]$SPACE\(.*\)$SPACE\$|\1$FS\2$FS\3|p" |

	awk -F"$FS" '{
		indent = length($1)/2;
		if (length($2) == 0) { conj[indent]="+";} else {conj[indent]="";}
		vname[indent] = $2;
		for (i in vname) {if (i > indent) {delete vname[i]}}
			if (length($3) > 0) {
				vn=""; for (i=0; i<indent; i++) {vn=(vn)(vname[i])(".")}
				printf("%s%s%s%s=%s\n", "'"$prefix"'",vn, $2, conj[indent-1],$3);
			}
		}' |

	sed -e 's/_=/+=/g' |

	awk 'BEGIN {
			FS="=";
			OFS="="
		}
		/(-|\.).*=/ {
			gsub("-|\\.", ".", $1)
		}
		{ print }' |

	grep "$KEY=" |

	cut -d'=' -f2
)

if [ -z $RESULT ]; then
	>&2 echo "Key $KEY not found"
	exit 1
else
	echo $RESULT
fi
