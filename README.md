**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-19 17:57:21 UTC（2026-03-20 01:57:21 UTC+8）

**代理总数：65**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1629ms | ✓ 1880ms | ✓ 1342ms | ✓ 1199ms | ✓ 1155ms | http |
| 174.138.24.77:1080 | ✓ 1708ms | 否 | 否 | ✓ 1253ms | ✓ 1013ms | http |
| 113.160.132.26:8080 | ✓ 1796ms | 否 | ✓ 1844ms | ✓ 1464ms | ✓ 1145ms | http |
| 45.167.124.52:8080 | ✓ 1108ms | 否 | ✓ 763ms | 否 | ✓ 1470ms | http |
| 45.125.67.37:443 | 否 | 否 | ✓ 1652ms | ✓ 1660ms | ✓ 1812ms | http |
| 85.198.96.242:3128 | ✓ 1159ms | 否 | ✓ 946ms | 否 | ✓ 1316ms | http |
| 38.34.179.63:8448 | ✓ 1299ms | 否 | ✓ 1125ms | ✓ 995ms | ✓ 1351ms | http |
| 91.238.104.171:2023 | ✓ 1155ms | 否 | ✓ 1834ms | 否 | ✓ 1674ms | http |
| 35.225.22.61:80 | ✓ 1034ms | ✓ 1858ms | ✓ 1056ms | ✓ 1236ms | 否 | http |
| 223.16.170.103:80 | 否 | 否 | ✓ 1600ms | ✓ 1385ms | ✓ 1330ms | http |
| 120.92.212.16:8890 | ✓ 1147ms | ✓ 1469ms | 否 | 否 | ✓ 1145ms | http |
| 137.220.150.152:6005 | ✓ 1098ms | 否 | 否 | ✓ 1666ms | ✓ 1324ms | http |
| 147.161.239.240:8800 | ✓ 1674ms | ✓ 1587ms | ✓ 1141ms | ✓ 1489ms | ✓ 1179ms | http |
| 38.34.179.29:8452 | ✓ 919ms | ✓ 1722ms | ✓ 1299ms | ✓ 976ms | ✓ 1304ms | http |
| 120.92.212.16:7890 | ✓ 1105ms | ✓ 1434ms | ✓ 1200ms | ✓ 1435ms | ✓ 1162ms | http |
| 101.43.127.100:8877 | ✓ 1087ms | ✓ 1339ms | 否 | ✓ 1283ms | ✓ 1006ms | http |
| 38.34.179.161:8450 | ✓ 1153ms | ✓ 1049ms | 否 | ✓ 1411ms | ✓ 836ms | http |
| 38.145.218.232:8448 | ✓ 1153ms | ✓ 1112ms | ✓ 992ms | 否 | ✓ 885ms | http |
| 38.145.218.212:8448 | ✓ 1166ms | ✓ 1163ms | ✓ 927ms | 否 | ✓ 901ms | http |
| 38.145.208.172:8448 | ✓ 1607ms | 否 | ✓ 1286ms | ✓ 1204ms | ✓ 1008ms | http |
| 38.145.208.247:8444 | ✓ 1610ms | ✓ 1266ms | 否 | ✓ 1323ms | ✓ 1116ms | http |
| 219.117.204.211:7799 | ✓ 1576ms | 否 | 否 | ✓ 1171ms | ✓ 960ms | http |
| 1.231.81.166:3128 | ✓ 1606ms | ✓ 1453ms | ✓ 1982ms | 否 | ✓ 1403ms | http |
| 38.34.179.32:8450 | ✓ 977ms | 否 | ✓ 954ms | ✓ 1114ms | ✓ 835ms | http |
| 38.34.179.175:8448 | ✓ 994ms | ✓ 1781ms | ✓ 1332ms | ✓ 931ms | ✓ 996ms | http |
| 38.145.220.37:8452 | ✓ 1289ms | ✓ 1279ms | ✓ 1062ms | 否 | ✓ 765ms | http |
| 38.145.220.15:8450 | ✓ 1283ms | ✓ 1279ms | ✓ 1067ms | 否 | ✓ 742ms | http |
| 38.34.179.176:8450 | ✓ 1288ms | 否 | ✓ 909ms | ✓ 959ms | ✓ 1130ms | http |
| 115.231.181.40:8128 | ✓ 1929ms | 否 | ✓ 1043ms | ✓ 1349ms | 否 | http |
| 45.136.198.40:3128 | ✓ 1441ms | ✓ 1632ms | ✓ 947ms | ✓ 1820ms | ✓ 1294ms | http |
| 137.220.151.110:6005 | ✓ 941ms | 否 | ✓ 1113ms | ✓ 1280ms | ✓ 1042ms | http |
| 59.46.216.131:30001 | ✓ 1161ms | 否 | 否 | ✓ 1637ms | ✓ 1220ms | http |
| 137.220.150.170:6005 | ✓ 1334ms | 否 | ✓ 1767ms | ✓ 1458ms | ✓ 1058ms | http |
| 38.145.220.198:8448 | ✓ 1019ms | 否 | ✓ 948ms | ✓ 1537ms | 否 | http |
| 133.242.138.34:8100 | ✓ 1818ms | ✓ 1412ms | 否 | 否 | ✓ 1939ms | http |
| 185.115.74.185:8080 | ✓ 1278ms | ✓ 1997ms | ✓ 1591ms | 否 | 否 | http |
| 38.34.179.35:8445 | ✓ 1228ms | ✓ 1523ms | ✓ 363ms | ✓ 917ms | ✓ 794ms | http |
| 38.145.208.181:8445 | ✓ 1969ms | ✓ 930ms | ✓ 616ms | 否 | ✓ 1302ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1667ms | ✓ 1530ms | ✓ 1416ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1393ms | ✓ 1117ms | ✓ 1244ms | ✓ 870ms | http |
| 38.145.208.167:8450 | ✓ 822ms | ✓ 1764ms | ✓ 416ms | ✓ 943ms | ✓ 1335ms | http |
| 38.145.208.229:8450 | ✓ 1813ms | ✓ 1876ms | ✓ 1652ms | ✓ 958ms | ✓ 1705ms | http |
| 38.145.203.132:8450 | ✓ 1813ms | ✓ 1249ms | ✓ 1523ms | ✓ 1647ms | ✓ 1746ms | http |
| 45.136.130.171:8445 | ✓ 1466ms | ✓ 941ms | ✓ 473ms | ✓ 1654ms | ✓ 1559ms | http |
| 86.53.183.16:1080 | ✓ 1009ms | 否 | ✓ 1183ms | 否 | ✓ 1331ms | http |
| 204.48.31.203:80 | ✓ 107ms | ✓ 1369ms | ✓ 1415ms | ✓ 1585ms | 否 | http |
| 106.117.208.101:7890 | ✓ 1328ms | ✓ 1630ms | ✓ 1346ms | ✓ 1561ms | ✓ 1313ms | http |
| 137.220.150.104:6005 | ✓ 935ms | 否 | ✓ 956ms | ✓ 1574ms | ✓ 1045ms | http |
| 20.27.15.49:8561 | ✓ 799ms | ✓ 1833ms | ✓ 1345ms | ✓ 1856ms | ✓ 1953ms | http |
| 20.210.76.175:8561 | ✓ 841ms | 否 | ✓ 1194ms | ✓ 1841ms | ✓ 1949ms | http |
| 20.210.76.104:8561 | ✓ 812ms | ✓ 1144ms | ✓ 1336ms | ✓ 1879ms | ✓ 1924ms | http |
| 20.210.76.178:8561 | ✓ 781ms | ✓ 1779ms | ✓ 1381ms | ✓ 1882ms | ✓ 1675ms | http |
| 3.99.169.21:42677 | ✓ 1929ms | 否 | ✓ 1595ms | 否 | ✓ 1966ms | http |
| 88.80.150.82:8080 | ✓ 1172ms | ✓ 1923ms | ✓ 1982ms | 否 | ✓ 1160ms | https |
| 45.168.238.193:8443 | 否 | ✓ 1741ms | ✓ 270ms | ✓ 1212ms | ✓ 907ms | http |
| 223.16.170.103:3128 | ✓ 1322ms | 否 | ✓ 1245ms | ✓ 1319ms | ✓ 1334ms | http |
| 103.39.51.190:8080 | ✓ 1914ms | 否 | 否 | ✓ 1919ms | ✓ 1945ms | http |
| 103.113.70.189:1081 | ✓ 178ms | ✓ 865ms | 否 | ✓ 1172ms | ✓ 687ms | http |
| 45.136.131.47:8449 | ✓ 1136ms | ✓ 1949ms | ✓ 1178ms | 否 | 否 | http |
| 121.230.8.220:1080 | ✓ 1442ms | ✓ 1684ms | ✓ 1385ms | ✓ 1882ms | ✓ 1603ms | http |
| 54.37.72.89:80 | ✓ 789ms | ✓ 1804ms | 否 | 否 | ✓ 1985ms | http |
| 61.52.131.172:8443 | ✓ 1009ms | ✓ 1368ms | ✓ 1095ms | ✓ 1390ms | ✓ 1120ms | http |
| 103.84.95.54:7890 | ✓ 1464ms | 否 | ✓ 889ms | 否 | ✓ 1220ms | http |
| 14.225.212.37:7890 | 否 | ✓ 1545ms | 否 | ✓ 1519ms | ✓ 1300ms | http |
| 180.125.216.109:8118 | 否 | ✓ 1256ms | ✓ 1034ms | 否 | ✓ 1314ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
