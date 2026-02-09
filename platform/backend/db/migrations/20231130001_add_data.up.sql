UPDATE "public"."agent_game_icon_list"
  SET "icon_list" = (SELECT jsonb_agg(jsonb_set(elems, '{push}', '0')) FROM jsonb_array_elements(icon_list) elems);