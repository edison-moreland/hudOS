-- windows = {}
-- windows.clock = {}
-- windows.clock.output = "DP-1"
-- windows.clock.width = 461
-- windows.clock.height = 112
-- windows.clock.x = 0
-- windows.clock.y = 0

require('windows')

active_windows={}
function position_window(view, output, window)
    view:move(output.x+window.x, output.y+window.y)
    view:resize(window.width, window.height)
end

function new_window(view, output, window)
    position_window(view, output, window) 

    local app_id = view:app_id()
    local x, y = view:pos()
    active_windows[app_id] = {}
    active_windows[app_id].x = x 
    active_windows[app_id].y = y
    view:on("destroy", function(view)
        active_windows[view:app_id()] = nil
    end)
end

kiwmi:on("view", function(view)
    print("view "..view:app_id())
    local window = windows[view:app_id()]
    if window == nil then
        print("window " .. view:app_id() .. " does not exist!")
        view:close()
        return
    end

    print("view "..view:app_id() .." using ouput "..window.output)
    local output = available_outputs[window.output]
    if output == nil then
        print("output " .. window.output .. " is not connected!")
        view:close()
        return
    end

    new_window(view, output, window)
end)

available_outputs={}
kiwmi:on("output", function(output)
    local name = output:name()
    available_outputs[name] = {}
    if name == "DSI-1" then
        -- Phone screen
        output:move(0, 0)
        available_outputs[name].x = 0
        available_outputs[name].y = 0
    else
        -- Glasses
        output:move(720, 0)
        available_outputs[name].x = 720
        available_outputs[name].y = 0
    end

    output:on("destroy", function(output)
        print("destroy "..output:name())
        available_outputs[output:name()] = nil
    end)
end)