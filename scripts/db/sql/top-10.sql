SELECT 
    word, 
    COUNT(*) as frequency
FROM 
    (SELECT unnest(words) as word FROM languages) as subquery
GROUP BY 
    word
ORDER BY 
    frequency DESC
LIMIT 
    10;