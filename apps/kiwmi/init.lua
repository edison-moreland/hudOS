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
        active_windows[view:app_id()] = view:pos()
    end)
end

function reposition_all_windows()
    -- Reposition all windows on output
    print("repositioning all! windows")

    for app_id in pairs(active_windows) do
        print("repositioning "..app_id)

        app_old_pos = active_windows[app_id]
        local view = kiwmi:view_at(app_old_pos.x, app_old_pos.y)
        if view == nil then
            print("THIS SHOULD NOT HAPPEN :(")
            view:close() -- error on purpose
        end

        local window = windows[app_id]
        local output = available_outputs[window.output]
        if output == nil then
            print("output " .. window.output .. " is not connected!")
            view:close()
            return
        end


        position_window(view, output, window) 
    end
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
    local x, y = output:pos() 

    available_outputs[name] = {} 
    available_outputs[name].x = x
    available_outputs[name].y = y

    output:on("destroy", function(output)
        print("destroy "..output:name())
        available_outputs[output:name()] = nil

        -- This is pretty hacky. The only time an output is destoryed is when the headset is being disconnected
        -- Everything would work correctly without the following code, if we had some way to know 
        -- that output "DSI-1"'s position had changed. The only reason this isnt a problem when the headset 
        -- connects, is that the compositor restarts whenever the headset is connected. (Which is another hack)
        -- This could probably be fixed by making the phone screen(DSI-1) the main output
        available_outputs["DSI-1"].x = 0
        available_outputs["DSI-1"].y = 0

        reposition_all_windows()
    end)
end)