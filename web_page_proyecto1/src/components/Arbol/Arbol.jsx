import React from 'react';
import './Arbol.css';
import axios from 'axios';
import Tree from '@naisutech/react-tree'
export default class Arbol extends React.Component {
  constructor(props) {
    super(props);
    this.state = { Procesos: [] };

    axios.get(`http://18.204.15.140:8080/treeprocess`)
      .then(res => {
        let arbol = res.data.Arbol;
        arbol.map((proc) => {
          proc = this.NoCapitalLetter(proc)
          return null;
        });
        //console.log(arbol)
        this.setState({ Procesos: arbol });
      })

  }

  NoCapitalLetter(proc) {
    let obj = proc;
    obj.id = parseInt(obj.Id);
    obj.parentId = parseInt(obj.ParentId);
    obj.items = obj.Items;
    obj.label = "(" + obj.id + ") " + obj.Label;
    delete obj.Id;
    delete obj.Label;
    delete obj.ParentId;
    delete obj.Items;
    if (obj.parentId === 0) {
      obj.parentId = null
    }
    obj.items.map((hijo) => {
      hijo = this.NoCapitalLetter(hijo)
      return null;
    });


    return obj
  }

  componentDidMount() {
    axios.get(`http://18.204.15.140:8080/treeprocess`)
      .then(res => {
        let arbol = res.data.Arbol;
        arbol.map((proc) => {
          proc = this.NoCapitalLetter(proc)
          if (proc.parentId === "0") {
            proc.parentId = null
          }
          return null;
        });
        let data = [
          {
            "id": 12345678,
            "parentId": null,
            "label": "My parent node",
            "items": [
              {
                "id": 87654321,
                "label": "My file",
                "parentId": 12345678,
                "items": []
              }
            ]
          },
          {
            "id": 56789012,
            "parentId": 12345678,
            "label": "My child node",
            "items": []
          }
        ];
        console.log(data)
        this.setState({ Procesos: arbol });
      })
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }


  render() {
    return (

      <div>
        <Tree
          nodes={this.state.Procesos} // see data format
          //onSelect={(a) => { a }} // fired every click of node or leaf with selected item as argument
          //        size={"full"} // full (default), half, narrow
          id="tree"
        />
      </div>
    );
  }
}

