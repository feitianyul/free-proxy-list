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

最后更新：2026-03-10 19:53:09 UTC（2026-03-11 03:53:09 UTC+8）

**代理总数：80**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.140.147.82:1081 | ✓ 422ms | ✓ 1493ms | ✓ 1233ms | 否 | ✓ 1132ms | http |
| 1.231.81.166:3128 | ✓ 1672ms | ✓ 1082ms | ✓ 946ms | ✓ 1162ms | ✓ 883ms | http |
| 222.184.48.236:22222 | ✓ 1211ms | 否 | ✓ 1327ms | 否 | ✓ 1477ms | http |
| 45.136.131.47:8443 | ✓ 756ms | ✓ 970ms | ✓ 730ms | ✓ 876ms | ✓ 663ms | http |
| 165.227.5.10:8888 | ✓ 432ms | 否 | ✓ 684ms | ✓ 1409ms | 否 | http |
| 190.9.109.198:999 | ✓ 699ms | ✓ 1425ms | ✓ 1051ms | ✓ 1309ms | ✓ 1262ms | http |
| 45.136.130.223:8443 | ✓ 335ms | ✓ 835ms | ✓ 777ms | 否 | 否 | http |
| 39.104.201.40:7890 | ✓ 1040ms | ✓ 1308ms | ✓ 1101ms | ✓ 1365ms | ✓ 1072ms | http |
| 81.70.169.194:80 | ✓ 1132ms | ✓ 1358ms | ✓ 1123ms | ✓ 1449ms | ✓ 1096ms | http |
| 101.43.255.96:80 | ✓ 1169ms | ✓ 1478ms | ✓ 1139ms | ✓ 1441ms | ✓ 1186ms | http |
| 91.107.141.42:8081 | ✓ 1058ms | 否 | ✓ 1383ms | 否 | ✓ 1935ms | http |
| 158.69.185.37:3129 | ✓ 798ms | ✓ 1667ms | ✓ 1137ms | ✓ 1446ms | ✓ 1006ms | http |
| 185.191.236.162:3128 | 否 | 否 | ✓ 1829ms | ✓ 1639ms | ✓ 1067ms | http |
| 45.186.6.104:3128 | ✓ 1350ms | ✓ 1766ms | ✓ 1703ms | 否 | 否 | http |
| 162.240.154.26:3128 | ✓ 799ms | ✓ 1718ms | ✓ 328ms | ✓ 922ms | ✓ 1393ms | http |
| 121.230.9.148:1080 | ✓ 1195ms | ✓ 1645ms | ✓ 1226ms | ✓ 1725ms | ✓ 1175ms | http |
| 120.92.212.16:8890 | ✓ 1245ms | ✓ 1691ms | ✓ 1352ms | 否 | ✓ 1374ms | http |
| 222.184.48.252:22222 | ✓ 1044ms | ✓ 1294ms | ✓ 1039ms | ✓ 1327ms | 否 | http |
| 202.155.12.161:443 | ✓ 1831ms | 否 | 否 | ✓ 1151ms | ✓ 1109ms | http |
| 120.198.141.75:22222 | ✓ 1083ms | ✓ 1440ms | ✓ 1117ms | ✓ 1287ms | ✓ 1031ms | http |
| 138.124.53.25:7443 | ✓ 1496ms | ✓ 1652ms | ✓ 1187ms | ✓ 1786ms | ✓ 1299ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1347ms | ✓ 1099ms | 否 | ✓ 1092ms | http |
| 101.47.73.135:3128 | ✓ 1339ms | 否 | ✓ 1591ms | ✓ 1661ms | 否 | http |
| 35.225.22.61:80 | 否 | ✓ 1031ms | ✓ 259ms | ✓ 1078ms | ✓ 727ms | http |
| 45.140.147.82:1082 | ✓ 828ms | 否 | ✓ 1232ms | ✓ 1405ms | ✓ 1205ms | http |
| 45.136.198.40:3128 | ✓ 1475ms | ✓ 1834ms | ✓ 1647ms | 否 | ✓ 1844ms | http |
| 120.198.141.80:22222 | ✓ 1233ms | ✓ 1395ms | ✓ 1216ms | ✓ 1335ms | ✓ 1049ms | http |
| 103.113.70.189:1081 | ✓ 335ms | ✓ 1393ms | 否 | ✓ 1181ms | ✓ 825ms | http |
| 95.3.9.78:8080 | ✓ 1327ms | ✓ 1902ms | ✓ 968ms | ✓ 1618ms | ✓ 1251ms | http |
| 95.3.9.78:3128 | ✓ 1330ms | ✓ 1782ms | ✓ 1088ms | ✓ 1650ms | ✓ 1263ms | http |
| 34.96.238.40:8080 | ✓ 1024ms | 否 | ✓ 1267ms | ✓ 1098ms | 否 | http |
| 37.139.33.145:1080 | ✓ 890ms | 否 | ✓ 723ms | ✓ 1964ms | ✓ 1802ms | http |
| 185.109.21.39:3128 | 否 | 否 | ✓ 785ms | ✓ 1732ms | ✓ 1324ms | http |
| 183.249.5.105:22222 | 否 | ✓ 1268ms | ✓ 1035ms | ✓ 1115ms | ✓ 821ms | http |
| 67.169.98.211:443 | ✓ 1285ms | 否 | 否 | ✓ 1629ms | ✓ 1391ms | http |
| 113.177.131.2:3128 | ✓ 1282ms | ✓ 1807ms | ✓ 1634ms | ✓ 1699ms | ✓ 1247ms | http |
| 162.248.165.72:1080 | ✓ 1034ms | 否 | ✓ 1848ms | ✓ 1610ms | 否 | http |
| 103.46.8.102:8080 | ✓ 1684ms | 否 | ✓ 1805ms | 否 | ✓ 1661ms | http |
| 107.178.115.140:3128 | ✓ 1253ms | ✓ 1624ms | ✓ 1170ms | 否 | 否 | http |
| 117.159.239.42:22222 | ✓ 957ms | ✓ 1252ms | ✓ 944ms | ✓ 1243ms | ✓ 937ms | http |
| 106.14.203.63:3333 | ✓ 907ms | 否 | ✓ 1546ms | ✓ 1238ms | ✓ 983ms | http |
| 180.125.216.109:8118 | 否 | ✓ 1367ms | ✓ 1140ms | ✓ 1382ms | 否 | http |
| 117.159.239.49:22222 | ✓ 961ms | ✓ 1161ms | ✓ 896ms | ✓ 1213ms | ✓ 955ms | http |
| 190.212.131.238:3128 | ✓ 1036ms | 否 | 否 | ✓ 1925ms | ✓ 1524ms | http |
| 120.238.159.229:22222 | ✓ 1018ms | ✓ 1307ms | ✓ 1080ms | ✓ 1254ms | ✓ 974ms | http |
| 34.101.184.164:3128 | ✓ 1611ms | 否 | ✓ 1984ms | ✓ 1549ms | ✓ 1228ms | http |
| 168.235.110.63:3128 | ✓ 329ms | 否 | ✓ 809ms | ✓ 1406ms | ✓ 1217ms | http |
| 128.199.114.189:9090 | ✓ 1731ms | 否 | ✓ 1351ms | ✓ 1596ms | ✓ 1408ms | http |
| 103.35.188.243:3128 | ✓ 326ms | ✓ 1977ms | 否 | 否 | ✓ 1024ms | http |
| 120.240.29.51:22222 | ✓ 1016ms | ✓ 1375ms | 否 | ✓ 1259ms | 否 | http |
| 152.42.213.210:8080 | ✓ 1984ms | 否 | ✓ 1401ms | ✓ 1215ms | ✓ 1171ms | http |
| 205.209.118.30:3138 | ✓ 142ms | ✓ 968ms | ✓ 363ms | ✓ 1156ms | ✓ 894ms | http |
| 222.184.48.235:22222 | 否 | ✓ 1420ms | 否 | ✓ 1250ms | ✓ 1067ms | http |
| 120.238.159.230:22222 | 否 | 否 | ✓ 1138ms | ✓ 1236ms | ✓ 1039ms | http |
| 59.46.216.131:30001 | ✓ 1187ms | ✓ 1834ms | ✓ 1181ms | ✓ 1479ms | ✓ 1116ms | http |
| 222.184.48.248:22222 | ✓ 1303ms | ✓ 1544ms | ✓ 1265ms | ✓ 1546ms | ✓ 1133ms | http |
| 103.39.51.190:8080 | ✓ 1875ms | 否 | 否 | ✓ 1529ms | ✓ 1519ms | http |
| 113.59.32.162:22222 | ✓ 1118ms | ✓ 1453ms | ✓ 1116ms | ✓ 1371ms | ✓ 1073ms | http |
| 115.231.181.40:8128 | ✓ 1500ms | ✓ 1686ms | ✓ 1858ms | 否 | 否 | http |
| 45.136.130.188:8443 | ✓ 800ms | ✓ 1013ms | ✓ 1110ms | ✓ 1022ms | ✓ 650ms | http |
| 45.136.130.175:8443 | ✓ 801ms | ✓ 993ms | ✓ 1130ms | ✓ 1022ms | ✓ 640ms | http |
| 45.136.130.191:8443 | ✓ 800ms | ✓ 1015ms | ✓ 1107ms | ✓ 992ms | ✓ 694ms | http |
| 45.136.131.63:8443 | ✓ 800ms | ✓ 1198ms | ✓ 926ms | ✓ 965ms | ✓ 776ms | http |
| 46.183.25.8:443 | ✓ 1575ms | 否 | 否 | ✓ 1166ms | ✓ 1307ms | http |
| 120.238.159.228:22222 | ✓ 984ms | ✓ 1362ms | ✓ 1048ms | ✓ 1260ms | ✓ 1006ms | http |
| 120.198.141.84:22222 | ✓ 1107ms | ✓ 1435ms | ✓ 1008ms | ✓ 1225ms | ✓ 1019ms | http |
| 159.223.42.219:3128 | ✓ 1510ms | 否 | ✓ 998ms | ✓ 1242ms | ✓ 966ms | http |
| 152.70.98.46:8888 | ✓ 1926ms | ✓ 1522ms | ✓ 1681ms | ✓ 1409ms | ✓ 1265ms | http |
| 222.184.48.251:22222 | ✓ 1094ms | ✓ 1337ms | ✓ 1178ms | ✓ 1257ms | ✓ 933ms | http |
| 101.36.224.130:7897 | ✓ 1033ms | ✓ 1310ms | ✓ 1054ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 1068ms | ✓ 1339ms | ✓ 997ms | ✓ 1270ms | ✓ 1019ms | http |
| 45.140.147.155:1082 | ✓ 771ms | 否 | ✓ 1440ms | ✓ 1635ms | ✓ 1274ms | http |
| 5.101.0.233:3128 | ✓ 728ms | 否 | ✓ 1131ms | 否 | ✓ 1854ms | http |
| 194.213.18.200:443 | ✓ 1928ms | ✓ 1902ms | ✓ 1855ms | 否 | 否 | http |
| 178.236.245.59:3128 | ✓ 1249ms | 否 | 否 | ✓ 1888ms | ✓ 1301ms | http |
| 178.236.245.17:3128 | ✓ 1220ms | 否 | ✓ 1984ms | ✓ 1815ms | ✓ 1256ms | http |
| 138.124.90.140:1080 | ✓ 490ms | 否 | ✓ 1551ms | ✓ 1583ms | 否 | http |
| 117.159.239.54:22222 | ✓ 904ms | ✓ 1204ms | ✓ 969ms | ✓ 1221ms | ✓ 975ms | http |
| 83.219.250.8:62920 | ✓ 1339ms | 否 | ✓ 1834ms | 否 | ✓ 1591ms | http |
| 120.240.35.177:22222 | ✓ 1668ms | 否 | 否 | ✓ 1461ms | ✓ 1182ms | http |

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
