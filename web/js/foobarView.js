const PropTypes = React.PropTypes;
import { fooBarToggleShowValues } from "./actions";

Array.prototype.flatMap = function(lambda) {
    return Array.prototype.concat.apply([], this.map(lambda));
};

const MainRow = ({ name, id, onClick, showValues}) => {
         var span = showValues ? <span className="glyphicon glyphicon-zoom-out"></span> : <span className="glyphicon glyphicon-zoom-in"></span>
         return (
              <tr key={id}>
                <td>{name}</td>
                <td>{id}</td>
                <td><button type="button" className="btn btn-default btn-sm" onClick={() => onClick(id)}>
                        {span}
                     </button>
                </td>
              </tr>
          )
}

const ValueRow = ({ name, id }) => {
         return (
              <tr key={id}>
                <td>{name}</td>
                <td>{id}</td>
              </tr>
          )
}

export const FooBarTable = ({ lines, onClick }) => {
    let rows = lines.flatMap((line) => {
            let main =  <MainRow name={line.name}  id={line.id} key={line.id} onClick={onClick} showValues={line.showValues} />
            let values = line.showValues ? line.values.map(v => <ValueRow name={v} id={line.id + "." + v} />) : [];
            return [main, ...values]
        });

    return (
      <table className="table table-hover table-condensed table-responsive">
        <thead>
          <tr>
            <th>Name</th>
            <th>Id</th>
            <th>Show Values</th>
          </tr>
        </thead>
        <tbody>{rows}</tbody>
      </table>
    );
}

//    This is the model
//     [
//    {"name": "name1", "id": "1" "showValues": false "values": ["a", "b", "c"]},
//    {"name": "name2", "id": "2" "showValues": false "values": ["a", "b", "c"]},
//    {"name": "name3", "id": "3" "showValues": false "values": ["a", "b", "c"]}
//      ]

FooBarTable.propTypes = {
  lines : PropTypes.arrayOf(
           PropTypes.shape({name : PropTypes.string.isRequired,
                            id : PropTypes.string.isRequired,
                            showValues : PropTypes.bool.isRequired,
                            values : PropTypes.arrayOf(PropTypes.string)})),
  onClick: PropTypes.func.isRequired
};



const { Component } = React;

const mapStateToProps = (state) => {
    return {
        lines : state.fooBar
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        onClick : (id) =>  { dispatch(fooBarToggleShowValues(id)) }
    }
}

const { connect } = ReactRedux;


export const FooBarTableComponent = connect(mapStateToProps, mapDispatchToProps)(FooBarTable);