-- windows = {}
-- windows.clock = {}
-- windows.clock.output = "DP-1"
-- windows.clock.width = 461
-- windows.clock.height = 112
-- windows.clock.x = 0
-- windows.clock.y = 0

require('windows')

kiwmi:on("view", function(view)
    local window = windows[view:app_id()]
    if window == nil then
        print("window " .. view:app_id() .. " does not exist!")
        view:close()
        return
    end

    if window.output ~= "DSI-1" then
        print("warning: multiple outputs not supported yet")
    end

    print("positioning " .. view:app_id())
    view:move(window.x, window.y)
    view:resize(window.width, window.height)
end)

kiwmi:on("output", function(output)
    print("New output!", output:name())
end)