package db

type seedContentData struct {
	Content  string
	Sections []string
}

func GetSeedContentData() []seedContentData {
	return []seedContentData{
		{
			Content: "<header><h1 id=job-title>Junior Software Developer</h1><p id=job-location>Location: [City, State]<p id=job-type>Type: [Full-time/Part-time]</header><section id=job-summary><h2>Job Summary</h2><p>We are seeking a motivated Junior Software Developer to join our team. This is an excellent opportunity for a recent graduate or someone with limited experience to gain hands-on experience in software development. As a Junior Software Developer, you will work under the guidance of senior developers to assist in the design, development, and testing of software applications.</section><section id=job-responsibilities><h2>Responsibilities</h2><ul><li>Assist in the design and development of software solutions.<li>Write clean, maintainable, and efficient code.<li>Participate in code reviews and provide feedback to peers.<li>Debug and troubleshoot software issues.<li>Collaborate with cross-functional teams to deliver high-quality products.</ul></section><section id=job-qualifications><h2>Qualifications</h2><ul><li>Bachelor's degree in Computer Science or related field.<li>Knowledge of programming languages such as Java, C#, or Python.<li>Understanding of basic software development concepts.<li>Strong problem-solving and analytical skills.<li>Excellent communication and teamwork abilities.</ul></section><section id=job-tech-skills><h2>Tech Skills Required</h2><ul><li>Java<li>C#<li>Python<li>Basic understanding of HTML, CSS, and JavaScript<li>Version Control Systems (e.g., Git)</ul></section><section id=job-preferred-qualifications><h2>Preferred Qualifications</h2><ul><li>Experience with database technologies (e.g., SQL, MongoDB).<li>Knowledge of software development methodologies (e.g., Agile, Scrum).<li>Experience with web development frameworks (e.g., Spring, Django).</ul></section><section id=job-benefits><h2>Benefits</h2><ul><li>Competitive salary<li>Healthcare coverage<li>401(k) retirement plan<li>Paid time off<li>Mentorship opportunities<li>Professional development support</ul></section><footer><p>For inquiries, please contact [Contact Information].</footer>",
			Sections: []string{"Job Title#job-title",
				"Job Location#job-location",
				"Job Type#job-type",
				"Job Summary#job-summary",
				"Responsibilities#job-responsibilities",
				"Qualifications#job-qualifications",
				"Tech Skills Required#job-tech-skills",
				"Preferred Qualifications#job-preferred-qualifications",
				"Benefits#job-benefits"},
		},
		{
			Content: "<header><h1 id=job-title>Mid-Level Software Engineer</h1><p id=job-location>Location: [City, State]<p id=job-type>Type: [Full-time/Part-time]</header><section id=job-summary><h2>Job Summary</h2><p>We are seeking an experienced Mid-Level Software Engineer to join our team. The ideal candidate will have a proven track record of delivering high-quality software solutions. As a Mid-Level Software Engineer, you will be responsible for designing, implementing, and maintaining complex software systems.</section><section id=job-responsibilities><h2>Responsibilities</h2><ul><li>Design and develop scalable and maintainable software solutions.<li>Lead and mentor junior developers in best practices and coding standards.<li>Collaborate with product managers and stakeholders to define project requirements.<li>Conduct code reviews and provide constructive feedback to peers.<li>Ensure code quality through unit testing and integration testing.</ul></section><section id=job-qualifications><h2>Qualifications</h2><ul><li>Bachelor's degree in Computer Science or related field.<li>4+ years of experience in software development.<li>Proficiency in programming languages such as Java, C#, or Python.<li>Strong understanding of software architecture and design patterns.<li>Excellent problem-solving and analytical skills.</ul></section><section id=job-tech-skills><h2>Tech Skills Required</h2><ul><li>Java<li>C#<li>Python<li>Spring<li>.NET<li>RESTful APIs<li>SQL</ul></section><section id=job-preferred-qualifications><h2>Preferred Qualifications</h2><ul><li>Experience with cloud platforms (e.g., AWS, Azure).<li>Knowledge of microservices architecture.<li>Experience with containerization technologies (e.g., Docker, Kubernetes).</ul></section><section id=job-benefits><h2>Benefits</h2><ul><li>Competitive salary<li>Healthcare coverage<li>401(k) retirement plan<li>Paid time off<li>Professional development opportunities<li>Flexible work arrangements</ul></section><footer><p>For inquiries, please contact [Contact Information].</footer>",
			Sections: []string{"Job Title#job-title",
				"Job Location#job-location",
				"Job Type#job-type",
				"Job Summary#job-summary",
				"Responsibilities#job-responsibilities",
				"Qualifications#job-qualifications",
				"Tech Skills Required#job-tech-skills",
				"Preferred Qualifications#job-preferred-qualifications",
				"Benefits#job-benefits",
			},
		},
		{
			Content: "<header><h1 id=job-title>Senior Frontend Developer</h1><p id=job-location>Location: San Francisco, CA<p id=job-type>Type: Full-time</header><section id=job-summary><h2>Job Summary</h2><p>We are seeking an experienced Senior Frontend Developer to join our team. The ideal candidate will have a passion for creating clean, efficient, and scalable user interfaces. As a Senior Frontend Developer, you will collaborate with cross-functional teams to design and implement user-facing features. You will also be responsible for optimizing applications for maximum speed and scalability.</section><section id=job-responsibilities><h2>Responsibilities</h2><ul><li>Develop new user-facing features using modern JavaScript frameworks (e.g., React, Angular, Vue.js).<li>Optimize applications for maximum speed and scalability.<li>Collaborate with back-end developers and designers to improve usability.<li>Participate in code reviews and provide constructive feedback to team members.<li>Stay up-to-date with emerging frontend technologies and best practices.</ul></section><section id=job-qualifications><h2>Qualifications</h2><ul><li>Bachelor's degree in Computer Science or a related field.<li>5+ years of experience in frontend development.<li>Proficiency in JavaScript, HTML, and CSS.<li>Experience with modern JavaScript frameworks such as React, Angular, or Vue.js.<li>Strong understanding of web performance optimization techniques.<li>Excellent communication and collaboration skills.</ul></section><section id=job-tech-skills><h2>Tech Skills Required</h2><ul><li>JavaScript<li>HTML5<li>CSS3<li>React.js<li>Angular.js<li>Vue.js<li>Webpack<li>Babel<li>Version control systems (e.g., Git)<li>Responsive design principles</ul></section><section id=job-preferred-qualifications><h2>Preferred Qualifications</h2><ul><li>Experience with backend technologies (e.g., Node.js, Express.js).<li>Knowledge of UI/UX design principles.<li>Contributions to open-source projects or personal coding projects.</ul></section><section id=job-benefits><h2>Benefits</h2><ul><li>Competitive salary<li>Comprehensive health benefits<li>Stock options<li>Flexible work hours<li>Remote work opportunities<li>Professional development stipend<li>Team-building activities</ul></section><footer><p>For inquiries, please contact [Contact Information].</footer>",
			Sections: []string{
				"Job Title#job-title",
				"Job Location#job-location",
				"Job Type#job-type",
				"Job Summary#job-summary",
				"Responsibilities#job-responsibilities",
				"Qualifications#job-qualifications",
				"Tech Skills Required#job-tech-skills",
				"Preferred Qualifications#job-preferred-qualifications",
				"Benefits#job-benefits",
			},
		},
	}
}
