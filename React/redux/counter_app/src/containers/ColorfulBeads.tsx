import { connect } from "react-redux";

import ColorfulBeads from "../components/ColorfulBeads";
import { CounterState } from "../reducer";

const mapStateToProps = (state: CounterState) => ({
  count: state.count,
});

export default connect(mapStateToProps)(ColorfulBeads);
