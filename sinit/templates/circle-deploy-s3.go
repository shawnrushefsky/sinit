package templates

import "text/template"

// CircleDeployS3 returns the circle ci deploy job and workflow for pushing a build artifact to S3
func CircleDeployS3() (t *template.Template) {
	const raw = `  
  deploy:
    docker:
      - image: circleci/python:2.7-jessie
    steps:
      - attach_workspace:
          # Must be absolute path or relative path from working_directory
          at: /tmp/workspace
      - run:
          name: Install awscli
          command: sudo pip install awscli
      - run:
          name: Deploy to S3
          command: aws s3 cp {{.Flags}} /tmp/workspace/{{.PersistPath}} s3://$BUCKET_NAME


workflows:
  version: 2
  build-deploy:
    jobs:
      - build
      - deploy:
          requires:
            - build
          filters:
            branches:
              only: master
`

	t, err := template.New("circle-deploy-s3.yml").Parse(raw)
	if err != nil {
		panic(err)
	}

	return
}
