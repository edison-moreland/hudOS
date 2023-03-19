opt.root_color = Color.new(0.0, 0.0, 0.0, 1.0)

local function on_create_container(container)
    container.geom.width = 461
    container.geom.height = 112
end
event:add_listener("on_create_container", on_create_container)

-- opt:add_mon_rule({
--     output="",
--     callback=function() print("output connected") end
-- })

-- opt:add_rule({
--     title="",
--     class="",
--     callback=function(con)
--         print("new_container")
--         print(con.app_id)
--     end
-- })