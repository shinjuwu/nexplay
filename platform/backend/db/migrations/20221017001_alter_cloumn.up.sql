ALTER TABLE "public"."marquee" 
  ALTER COLUMN "lang" TYPE varchar(10) USING "lang"::varchar(10);


UPDATE marquee SET lang='zh-tw' WHERE lang='0';
UPDATE marquee SET lang='zh-cn' WHERE lang='1';
UPDATE marquee SET lang='en-us' WHERE lang='2';
UPDATE marquee SET lang='vi-vn' WHERE lang='3';
UPDATE marquee SET lang='th-th' WHERE lang='4';