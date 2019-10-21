#!/bin/sh
# 테스트할 때 마다 명령어치기 힘들어서 미리 필요한 명령어를 입력해 놓은 bash파일.
kalena -add -title data2 -layer title -start 2019-11-05T22:08:50+09:00 -end 2019-11-10T23:10:20+09:00 -user user1
kalena -add -title data3 -layer title -start 2019-12-05T22:08:50+09:00 -end 2019-12-10T23:10:20+09:00 -user user1
kalena -add -title woong -layer se -start 2019-10-16T20:30:00+09:00 -end 2019-10-16T22:30:00+09:00 -user user1


kalena -add -title data2 -layer se -start 2019-11-05T22:08:50+09:00 -end 2019-11-10T23:10:20+09:00 -user bae
kalena -add -title park -layer se -start 2019-10-21T14:10:00+09:00 -end 2019-10-21T15:00:00+09:00 -user bae
kalena -add -title woong -layer se -start 2019-10-16T20:30:00+09:00 -end 2019-10-16T22:30:00+09:00 -user bae

