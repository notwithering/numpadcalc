#!/bin/bash

NAME="numpadcalc"
TMP_DIR="/tmp/$NAME"
FIFO="$TMP_DIR/daemon"

mkdir -p "$TMP_DIR"
rm -f "$FIFO"
mkfifo "$FIFO"

cat <<EOF > "$TMP_DIR/xbindkeysrc"
"echo '0' > $FIFO"
    m:0x0 + KP_0

"echo '1' > $FIFO"
    m:0x0 + KP_1

"echo '2' > $FIFO"
    m:0x0 + KP_2

"echo '3' > $FIFO"
    m:0x0 + KP_3

"echo '4' > $FIFO"
    m:0x0 + KP_4

"echo '5' > $FIFO"
    m:0x0 + KP_5

"echo '6' > $FIFO"
    m:0x0 + KP_6

"echo '7' > $FIFO"
    m:0x0 + KP_7

"echo '8' > $FIFO"
    m:0x0 + KP_8

"echo '9' > $FIFO"
    m:0x0 + KP_9

"echo '.' > $FIFO"
    m:0x0 + KP_Decimal

"echo '/' > $FIFO"
    m:0x0 + KP_Divide

"echo '*' > $FIFO"
    m:0x0 + KP_Multiply

"echo '-' > $FIFO"
    m:0x0 + KP_Subtract

"echo '+' > $FIFO"
    m:0x0 + KP_Add

"echo 'eval' > $FIFO"
    m:0x0 + KP_Enter
EOF

xbindkeys -n -f "$TMP_DIR/xbindkeysrc" &

XBINDS_PID=$!

cleanup() {
    kill $XBINDS_PID
    rm -f "$TMP_DIR/xbindkeysrc"
    rm -f "$FIFO"
}
trap cleanup EXIT

formula=""

while true; do
    if read line < "$FIFO"; then
        if [[ -z "$formula" ]]; then
            clipboard=$(xclip -o -sel clip)

            if [[ "$clipboard" =~ ^[0-9]+$ ]] && ! [[ "$line" =~ ^[0-9]+$ ]]; then
                formula="$clipboard"
            else
                if ! [[ "$line" =~ ^[0-9]+$ ]]; then
                    formula="0"
                fi
            fi
        fi

        case "$line" in
            "eval")
                result=$(echo "$formula" | bc)
                formula=""

                printf "$result" | xclip -sel clip
                ;;

            *)
                formula="$formula$line"
                ;;
        esac
    fi
done
