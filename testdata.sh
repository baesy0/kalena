#!/bin/sh
# 테스트할 때 마다 명령어치기 힘들어서 미리 필요한 명령어를 입력해 놓은 bash파일.
kalena -add -collection user1 -title data2 -layer title -start 2019-11-05T22:08:50+09:00 -end 2019-11-10T23:10:20+09:00
kalena -add -collection user1 -title data3 -layer title -start 2019-12-05T22:08:50+09:00 -end 2019-12-10T23:10:20+09:00
kalena -add -collection user1 -title woong -layer se -start 2019-10-16T20:30:00+09:00 -end 2019-10-16T22:30:00+09:00

# 월별 일정 검색용 데이터. 11월 검색 기준
# ------- 검색되면 안되는 데이터 ---------
# 11월 이전에 끝남.
kalena -add -collection bae -title park -layer se -start 2019-10-21T14:10:00+09:00 -end 2019-10-23T15:00:00+09:00
# 11월 이후에 시작함.
kalena -add -collection bae -title woong -layer se -start 2019-12-16T20:30:00+09:00 -end 2019-12-25T22:30:00+09:00

# ------- 검색 되어야 하는 데이터 --------
# 11월에 포함
kalena -add -collection bae -title data2 -layer se -start 2019-11-05T22:08:50+09:00 -end 2019-11-10T23:10:20+09:00
# 11월 이전에 시작해서 11월에 일정이 끝나는 경우.
kalena -add -collection bae -title woong -layer se -start 2019-10-16T20:30:00+09:00 -end 2019-11-25T22:30:00+09:00
# 11월에 시작해서 11월 이후에 일정이 끝나는 경우.
kalena -add -collection bae -title woong -layer se -start 2019-11-16T20:30:00+09:00 -end 2019-12-25T22:30:00+09:00


# RestAPI
curl -d "collection=woong&title=todo&layer=study&start=2019-10-10T10:10:10%2B09:00&end=2019-10-11T10:10:10%2B09:00" http://127.0.0.1/api/add