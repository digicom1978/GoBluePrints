{\rtf1\ansi\ansicpg1252\cocoartf1038\cocoasubrtf360
{\fonttbl\f0\fswiss\fcharset0 Helvetica;}
{\colortbl;\red255\green255\blue255;}
\margl1440\margr1440\vieww9000\viewh8400\viewkind0
\pard\tx566\tx1133\tx1700\tx2267\tx2834\tx3401\tx3968\tx4535\tx5102\tx5669\tx6236\tx6803\ql\qnatural\pardirnatural

\f0\fs24 \cf0 package main\
\
import (\
	"log"\
	"net/http"\
)\
\
func main() \{\
\
	http.HandleFunc ("/", func (w http.ResponseWriter, r *http.Request) \{\
		w.Write( []byte (`\
			<html>\
				<head>\
					<title>Chat</title>\
				</head>\
				<body>\
					Let's chat!\
				</body>\
			</html>\
		`))\
	\})\
	// start the web server\
	if err := http.ListenAndServe (":8080", nil); err := nil \{\
		log.Fatal ("ListenAndServe:", err)\
	\}\
\}}