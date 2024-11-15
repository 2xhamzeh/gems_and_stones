local M = {}

M.highscore = nil
M.currentscore = 0

function M.load_highscore()
	local filename = sys.get_save_file("sys_save_load", "highscore") 
	local data = sys.load(filename) 
	M.highscore = data.highscore or 0
	return M.highscore
end

function M.save_highscore(highscore)
	local filename = sys.get_save_file("sys_save_load", "highscore")
	sys.save(filename, { highscore = highscore })  
end

return M