# effective_mobile_music_info

To run project. It will create a database, make migration, build app and start app.
```bash
make run
```

To create song
```bash
curl --request POST \
  --url http://localhost:8080/api/v1/create \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/10.0.0' \
  --data '{
	"group": "ABBA",
	"song": "Dancing Queen"
}'
```

To update song
```bash
curl --request PUT \
  --url http://localhost:8080/api/v1/update \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/10.0.0' \
  --data '{
	"groupId": "1",
	"song": "Dancing Queen",
	"releaseDate": "16.08.1976",
	"text": "You can dance, you can jive, having the time of your life\nSee that girl, watch that scene, diggin'\'' the dancing queen\nFriday night and the lights are low,\nLooking out for the place to go,\nWhere they play the right music, getting in the swing.\nYou come in to look for a king.\nAnybody could be that guy,\nNight is young and the music'\''s high.\nWith a bit of rock music, everything is fine,\nYou'\''re in the mood for a dance.\nAnd when you get the chance...\nChorus:\nYou are the dancing queen, young and sweet, only seventeen.\nDancing queen, feel the beat from the tambourine.\nYou can dance, you can jive, having the time of your life.\nSee that girl, watch that scene, diggin'\'' the dancing queen.\nYou'\''re a teaser, you turn '\''em on,\nLeave them burning and then you'\''re gone.\nLooking out for another, anyone will do,\nYou'\''re in the mood for a dance.\nAnd when you get the chance...\nChorus\nDiggin'\'' the dancing queen.",
	"link": "https://www.youtube.com/watch?v=xFrGuyw1V8s"
}'
```

To get song
```bash
curl --request GET \
  --url 'http://localhost:8080/api/v1/info?group=ABBA&song=Dancing%20Queen' \
  --header 'User-Agent: insomnia/10.0.0'
```