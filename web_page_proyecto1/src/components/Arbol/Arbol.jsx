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
          return true;
        });
        //console.log(arbol)
        let b = arbol
        this.setState({ Procesos: b });
      })

  }

  NoCapitalLetter(proc) {
    let obj = proc;
    obj.id = parseInt(obj.Id);
    obj.label = "(" + obj.id + ") " + obj.Label;
    obj.parentId = parseInt(obj.ParentId);
    obj.items = obj.Items;

    delete obj.Id;
    delete obj.Label;
    delete obj.ParentId;
    delete obj.Items;
    if (obj.parentId === 0) {
      obj.parentId = null
    }
    if (obj.id == 899) {
      console.log(obj)
    }
    obj.items.map((hijo) => {
      hijo = this.NoCapitalLetter(hijo)
      return true;
    });


    return obj
  }

  componentDidMount() {
    this.interval = setInterval(() => {
      axios.get(`http://18.204.15.140:8080/treeprocess`)
        .then(res => {
          let arbol = res.data.Arbol;
          arbol.map((proc) => {
            proc = this.NoCapitalLetter(proc)
            return null;
          });
          let data = [
            {
              "id": 1,
              "parentId": null,
              "label": "My parent node",
              "items": [
                {
                  "id": 87654321,
                  "label": "My file",
                  "parentId": 1,
                  "items": []
                }
              ]
            },
            {
              "id": 56789012,
              "parentId": 1,
              "label": "My child node",
              "items": []
            }
            , {
              "id": 567890121,
              "parentId": 1,
              "label": "My child node",
              "items": []
            }
          ];
          //console.log(data)
          this.setState({ Procesos: arbol });
        })
    }, 5000)
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

