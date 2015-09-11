task :build do
	
		sh  "go get"
		sh  "go build"
end

namespace :run do

	task :prod do
		ENV['PORT']="8080"
		ENV['SLACKKEY']="xoxp-2223026996-2229561973-10411867873-e89cc1"
		sh  "go get"
		sh  "go build"
		sh  "./bilious-funicular"
	end
end