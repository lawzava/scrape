#!/usr/bin/env bash

CURDIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
SCRAPER=$CURDIR/bin/scrape.darwin-amd-64

# TSV file with url and keyname


# while getopts "n:" opt; do
 # case $opt in
#    n)
#        n_opt="$OPTARG"
#    ;;
#    p) p_out="$OPTARG"
#    ;;
#    \?) echo "Invalid option -$OPTARG" >&2
#    ;;
#  esac
#done

file=$1
n_opt=100000000;


#columnm of URL
#url_col=10;
url_col=6;
#column for key
key_col=7;


# Help documentation
function help_info {

cat <<helpdoc
scrape.sh file
    File = Tab seperated file
helpdoc
    exit 1;
}


##
## MAIN
##
if [[ ! -f "$file" ]]; then
    echo "Arg #1 must tab seperated file path"
    help_info
fi


i=0
while read line; do
#echo $line


#echo $(cut -d$'\t' -f 6 <<< "$line")
#exit;

    url=$(cut -d$'\t' -f ${url_col} <<< "$line")
    key=$(cut -d$'\t' -f ${key_col} <<< "$line")


    url_regex='(https?|ftp|file)://[-A-Za-z0-9\+&@#/%?=~_|!:,.;]*[-A-Za-z0-9\+&@#/%=~_|]'
    if [[ ${url} =~ ${url_regex} ]]; then
        #>&2 echo "URL: $url"

        [[ ${url} =~ ^([^:]+://([^/]+))/?(.*)$ ]] && server="${BASH_REMATCH[2]}"
        >&2 echo "SERVER: ${server}"
    else
        >&2 echo "Invalid URL: $url"
        continue
    fi


    # run scraper
    ("$SCRAPER" -w "$url") | while read email; do
        echo -e "${email}\t${key}\t${server}\t${url}";
    done;

    i=$[$i +1]
    if (( $i > $n_opt )); then
        exit 0;
    fi

done < "$file"
