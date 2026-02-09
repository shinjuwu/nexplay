UPDATE "public"."agent_game_icon_list"
  SET "icon_list" = (SELECT jsonb_agg(elems - 'push') FROM jsonb_array_elements(icon_list) elems);