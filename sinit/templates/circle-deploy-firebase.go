package templates

import "text/template"

// CircleDeployFirebase returns the circle ci deploy job and workflow for deploying a site to Google Firebase
func CircleDeployFirebase() (t *template.Template) {
	const raw = `  
  deploy:
    docker:
      - image: google/cloud-sdk
    steps:
      - attach_workspace:
          # Must be absolute path or relative path from working_directory
          at: /tmp/workspace

      - run:
          name: Create keyfile from env
          command: echo $GCLOUD_SERVICE_KEY >> /tmp/key.json

      - run:
          name: Set up gcloud
          command: gcloud auth activate-service-account --key-file=/tmp/key.json && gcloud --quiet config set project $GOOGLE_PROJECT_ID
    
      - run:
          name: Upload to Storage Bucket
          command: gsutil cp {{.Flags}} /tmp/workspace/{{.PersistPath}}/. gs://$BUCKET_NAME

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
              only: master`

	t, err := template.New("circle-deploy-firebase.yml").Parse(raw)
	if err != nil {
		panic(err)
	}

	return
}
