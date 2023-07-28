# removing following icons:
# BIconCalendar
# BIconCalendarFill
# BIconChevronBarLeft
# BIconChevronDoubleLeft
# BIconChevronDown
# BIconChevronLeft
# BIconChevronUp
# BIconCircleFill
# BIconClock
# BIconClockFill
# BIconDash
# BIconPlus
# BIconStar
# BIconStarFill
# BIconStarHalf
# BIconX

# Regex explanation
# '15,$' - ignore first 15 lines and go until end
# '\bBIconCalendar\b' - \b in front and back -> Word must match completely, __BIconCalendar, BIconCalendar__, __BIconCalendar__ would not match
# \bBIconCalendar\b\|\bBIconCalendarFill\b - pipe (' | ') -> OR (match BIConCalendar or bBIconCalendarFill)
# !d - (at the end) - delete line that matches criteria
sed -i -e '15,${/\bBIconCalendar\b\|\bBIconCalendarFill\b\|\bBIconChevronBarLeft\b\|\bBIconChevronDoubleLeft\b\|\bBIconChevronDown\b\|\bBIconChevronLeft\b\|\bBIconChevronUp\b\|\bBIconCircleFill\b\|\bBIconClock\b\|\bBIconClockFill\b\|\bBIconDash\b\|\BIconPersonFill\b\|\bBIconPlus\b\|\bBIconStar\b\|\bBIconStarFill\b\|\bBIconStarHalf\b\|\bBIconX\b/!d}' node_modules/bootstrap-vue/esm/icons/icons.js

# removing icons from plugin.js is just for silecing WARNING 
# ex. "export 'BIconXCircleFill' was not found in './icons'
# functionality for removing icons and reducing size is in icons.js
#
# !!! WARNING !!!
# this regex isn't working completely, as it sometimes cuts in the middle of line
# from line 60 to 630 is working safely, but it's not deleting all lines that contains icons
# that's why I'm not executing this script 
# sed -i -e '60,630{/\bBIconCalendar\b\|\bBIconCalendarFill\b\|\bBIconChevronBarLeft\b\|\bBIconChevronDoubleLeft\b\|\bBIconChevronDown\b\|\bBIconChevronLeft\b\|\bBIconChevronUp\b\|\bBIconCircleFill\b\|\bBIconClock\b\|\bBIconClockFill\b\|\bBIconDash\b\|\BIconPersonFill\b\|\bBIconPlus\b\|\bBIconStar\b\|\bBIconStarFill\b\|\bBIconStarHalf\b\|\bBIconX\b/!d}' node_modules/bootstrap-vue/esm/icons/plugin.js