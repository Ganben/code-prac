select
    weather.id as 'id'
from
    weather
        join
    weather w on datediff(weather.date, w.date) = 1
        and weather.Temperature > w.Temperature
;
