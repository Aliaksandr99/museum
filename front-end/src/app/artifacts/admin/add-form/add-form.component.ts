import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';
import {FormArray, FormControl, FormGroup} from '@angular/forms';

import {Artifact, ArtifactService} from '../../../shared/artifactService';

@Component({
    selector: 'app-add-form',
    templateUrl: './add-form.component.html',
})

export class AddFormComponent implements OnInit {

    private addForm: FormGroup;

    constructor(private artifactService: ArtifactService,
                private router: Router
    ) {
    }


    ngOnInit() {
        this.initForm()
    }

    private initForm() {
        this.addForm = new FormGroup({
            creator: new FormControl(''),
            artifact_style: new FormControl(''),
            date_exc: new FormControl(''),
            transferred_by: new FormControl(''),
            length: new FormControl(''),
            height: new FormControl(''),
            width: new FormControl(''),
            objectGroupMap: new FormGroup({})
        })
    }

    // TODO: objectGroupMap -> (fromGroup)objectGroup : (FormArray [Controls] (objects)) subObject
    addObjectGroup(objectGroupName: string) {
        let emptyFromArray = new FormArray([new FormControl(''),]);
        this.addForm.controls["objectGroupMap"].value
            .addControl(objectGroupName, emptyFromArray);
    }

    addArtifact() {
        console.log("hello from form");
        console.log(this.addForm.value);

        let result: Artifact = {
            creator: this.addForm.controls["creator"].value,
            artifact_style: this.addForm.controls["artifact_style"].value,
            date_exc: this.addForm.controls["date_exc"].value,
            transferred_by: this.addForm.controls["transferred_by"].value,
            artifact_measurement: {
                length: +this.addForm.controls["length"].value,
                width: +this.addForm.controls["width"].value,
                height: +this.addForm.controls["height"].value,
            },
        };

        //artifact.id = this.artifactService.getArtifact(this.artifactService.getArtifactNumber() - 1).id + 1;
        this.artifactService.add(result);

        this.router.navigate(['/artifact'])
    }

}
