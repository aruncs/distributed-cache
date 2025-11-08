import React from "react"
import { NodeDetails } from "../../components/node-details"
import {makeHttpGet} from "../../utils/http-helper"
import './style.scss'
class CacheDetails extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            nodes: [],
            monitoring: false
        }
        this.handleToggleMonitor = this.handleToggleMonitor.bind(this)
    }
    async componentDidMount() {

        
    }
    async componentDidUpdate(prevProp, prevState) {
        if (prevState.monitoring == false && this.state.monitoring == true) {
            this.timerHandle = setInterval(async ()=>{
                let nodes = await this.fetchServerDetails()
                this.setState({
                    nodes
                })
            }, 2000)
        } else if(prevState.monitoring == true && this.state.monitoring == false) {
            clearInterval(this.timerHandle)
        }
    }
    handleToggleMonitor() {
        let monitoring = this.state.monitoring
        this.setState({
            monitoring: !monitoring
        })
    }
    render() {
        
        let {nodes, monitoring} = this.state
        console.log("node: ", nodes)
        let switchLabel = monitoring ? "Turn Off" : "Turn On"

        return (
            <div className="cache-details">
                <div className={`toggle-monitor-button ${monitoring ? 'on' : 'off'}`} onClick={this.handleToggleMonitor}>{switchLabel}</div>
                <div className="node-list">
                    {nodes && nodes.map((node)=>{
                    return (<NodeDetails id={node.id} name={node.name} data={node.data}></NodeDetails>)
                })}
                </div>
            </div>
        )
    }

    async fetchServerDetails() {
        let nodes = await makeHttpGet("http://localhost:8080/nodes")
        let updatedNodes = []
        for(let i=0; i < nodes.length; i++){
            let node = nodes[i]
            let data = await makeHttpGet(node.address + "/data")
            console.log("data is ", data)
            node.data = data
            updatedNodes.push(node)
        }
        return updatedNodes
    }
}

export default CacheDetails