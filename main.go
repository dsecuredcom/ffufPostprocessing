package main

/*
ffuf -o /tmp/ffuf/resuls.json -od /tmp/ffuf/bodies -u 'https://www.damianstrobel.de/FUZZ' -w /home/damian/Pentesting/wordlists/filesdirs/BASIC-with-EXT.txt -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36' -t 15 -c -v -timeout 5 -D -e php,aspx,jsp,zip,rar,tar.gz,html,log,js,txt,sql,sql.gz,pem,csv -X GET -maxtime 1200 -maxtime-job 600 -rate 175 -mc 200,301,302,400,401,402,403,405,429,500,502,503,504 -fc 404 -st
*/

func main() {

}
