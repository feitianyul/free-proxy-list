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

最后更新：2026-04-30 10:59:03 UTC（2026-04-30 18:59:03 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.71.229.255:3128 | ✓ 1300ms | ✓ 1318ms | ✓ 785ms | ✓ 1389ms | ✓ 1273ms | http |
| 218.108.131.186:17890 | ✓ 1020ms | ✓ 1460ms | ✓ 1041ms | ✓ 1277ms | ✓ 1107ms | http |
| 115.231.181.40:8128 | ✓ 1213ms | ✓ 1429ms | 否 | 否 | ✓ 1217ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1310ms | ✓ 1560ms | ✓ 1206ms | http |
| 212.58.132.5:8888 | ✓ 1868ms | 否 | ✓ 1485ms | ✓ 1500ms | ✓ 1204ms | http |
| 132.226.235.199:1080 | ✓ 1325ms | 否 | 否 | ✓ 1439ms | ✓ 1312ms | http |
| 45.167.124.71:999 | ✓ 1288ms | 否 | ✓ 1496ms | ✓ 1708ms | ✓ 1617ms | http |
| 47.85.51.197:1080 | ✓ 816ms | 否 | ✓ 38ms | ✓ 976ms | ✓ 757ms | http |
| 157.0.142.246:10057 | ✓ 1165ms | 否 | ✓ 1173ms | 否 | ✓ 1239ms | http |
| 168.110.52.228:3128 | ✓ 1367ms | ✓ 1881ms | ✓ 1822ms | ✓ 1335ms | ✓ 1019ms | http |
| 120.92.212.16:8890 | ✓ 1195ms | 否 | ✓ 1344ms | ✓ 1573ms | 否 | http |
| 159.223.225.118:8888 | 否 | ✓ 1901ms | ✓ 1484ms | ✓ 1811ms | 否 | http |
| 103.70.114.149:3128 | ✓ 1756ms | 否 | ✓ 1873ms | ✓ 1829ms | ✓ 1804ms | http |
| 34.101.184.164:3128 | ✓ 971ms | 否 | ✓ 1054ms | ✓ 1548ms | ✓ 1109ms | http |
| 77.110.107.80:8080 | ✓ 462ms | ✓ 1657ms | ✓ 381ms | ✓ 1794ms | ✓ 1432ms | http |
| 38.92.10.152:57579 | ✓ 818ms | ✓ 1187ms | ✓ 1288ms | ✓ 1204ms | ✓ 759ms | http |
| 38.92.10.98:20058 | ✓ 829ms | ✓ 1495ms | ✓ 570ms | ✓ 1227ms | ✓ 1944ms | http |
| 94.72.109.214:8888 | ✓ 1096ms | 否 | ✓ 523ms | ✓ 1468ms | ✓ 1129ms | http |
| 8.154.21.175:3128 | ✓ 1080ms | ✓ 1240ms | ✓ 1046ms | ✓ 1310ms | ✓ 1089ms | http |
| 80.92.204.47:1081 | 否 | ✓ 1323ms | ✓ 389ms | 否 | ✓ 1072ms | http |
| 65.109.213.99:1080 | ✓ 982ms | 否 | ✓ 1916ms | 否 | ✓ 1980ms | http |
| 107.173.160.222:1080 | ✓ 1425ms | 否 | ✓ 1789ms | ✓ 1233ms | ✓ 1189ms | http |
| 194.87.26.83:3128 | ✓ 910ms | 否 | ✓ 1348ms | ✓ 1392ms | 否 | http |
| 206.206.126.177:2412 | 否 | 否 | ✓ 1340ms | ✓ 1894ms | ✓ 1449ms | http |
| 59.46.216.131:30001 | ✓ 1212ms | 否 | ✓ 1274ms | 否 | ✓ 1319ms | http |
| 46.101.95.183:8888 | ✓ 1143ms | 否 | ✓ 988ms | ✓ 1605ms | 否 | http |
| 20.164.75.153:8080 | ✓ 1739ms | 否 | ✓ 1591ms | 否 | ✓ 1978ms | http |
| 106.10.55.212:1121 | ✓ 1094ms | 否 | ✓ 1265ms | ✓ 1660ms | 否 | http |
| 130.61.174.200:1080 | ✓ 1274ms | 否 | ✓ 531ms | 否 | ✓ 936ms | http |
| 45.129.141.143:3128 | ✓ 664ms | 否 | ✓ 1663ms | ✓ 1907ms | ✓ 1526ms | http |
| 103.157.200.126:3128 | ✓ 1326ms | 否 | ✓ 1182ms | ✓ 1658ms | ✓ 1544ms | http |
| 1.231.81.166:3128 | ✓ 1647ms | ✓ 1559ms | ✓ 1130ms | 否 | 否 | http |
| 77.110.119.136:3128 | ✓ 1555ms | 否 | ✓ 507ms | ✓ 1409ms | ✓ 1070ms | http |
| 162.240.154.26:3128 | ✓ 1840ms | ✓ 1936ms | 否 | 否 | ✓ 1603ms | http |
| 154.64.232.35:8080 | 否 | ✓ 1391ms | 否 | ✓ 1489ms | ✓ 1068ms | http |
| 120.92.108.86:7890 | ✓ 1389ms | 否 | ✓ 1438ms | 否 | ✓ 1553ms | http |
| 43.133.44.89:8888 | ✓ 1095ms | 否 | 否 | ✓ 1255ms | ✓ 965ms | http |
| 77.110.116.224:3128 | ✓ 515ms | 否 | ✓ 1407ms | 否 | ✓ 1326ms | http |
| 91.217.81.131:1080 | ✓ 1504ms | ✓ 1745ms | 否 | 否 | ✓ 1731ms | http |
| 62.113.119.14:8080 | 否 | 否 | ✓ 1015ms | ✓ 1676ms | ✓ 1208ms | http |
| 113.176.92.71:3128 | ✓ 1775ms | ✓ 1718ms | ✓ 1609ms | ✓ 1716ms | 否 | http |
| 168.222.254.136:8888 | ✓ 1205ms | ✓ 1698ms | ✓ 1785ms | 否 | 否 | http |
| 210.223.44.230:3128 | ✓ 1741ms | ✓ 1933ms | 否 | ✓ 1159ms | 否 | http |
| 103.3.246.71:3128 | ✓ 1805ms | 否 | ✓ 1271ms | ✓ 1498ms | ✓ 1535ms | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 1030ms | ✓ 1495ms | ✓ 1298ms | http |
| 183.238.3.150:7897 | ✓ 1869ms | ✓ 1337ms | ✓ 1106ms | ✓ 1313ms | ✓ 1015ms | http |
| 3.101.133.120:80 | ✓ 895ms | ✓ 1415ms | ✓ 1615ms | ✓ 1418ms | ✓ 1371ms | http |
| 120.92.212.16:7890 | ✓ 1986ms | ✓ 1616ms | ✓ 1679ms | 否 | 否 | http |
| 8.209.238.110:47701 | ✓ 1107ms | 否 | ✓ 990ms | ✓ 1054ms | ✓ 846ms | http |
| 89.208.106.138:10808 | 否 | ✓ 1782ms | ✓ 962ms | 否 | ✓ 1145ms | http |
| 45.140.147.82:1081 | ✓ 1590ms | ✓ 1317ms | ✓ 978ms | 否 | 否 | http |
| 45.153.231.229:8080 | ✓ 1923ms | ✓ 1774ms | ✓ 1400ms | 否 | ✓ 1887ms | http |
| 118.113.244.193:1080 | ✓ 1409ms | ✓ 1946ms | ✓ 1529ms | 否 | ✓ 1479ms | http |
| 152.70.91.193:40000 | ✓ 1720ms | 否 | ✓ 1954ms | ✓ 1864ms | ✓ 1604ms | http |
| 64.188.67.154:1080 | ✓ 1555ms | 否 | 否 | ✓ 1629ms | ✓ 1260ms | http |
| 45.140.147.82:1082 | ✓ 441ms | ✓ 1397ms | ✓ 1463ms | ✓ 1384ms | 否 | http |
| 185.21.11.140:1080 | ✓ 1857ms | ✓ 1320ms | ✓ 771ms | ✓ 1764ms | ✓ 1752ms | http |
| 37.187.109.70:10111 | ✓ 1026ms | ✓ 1825ms | ✓ 1187ms | 否 | 否 | http |
| 121.130.177.28:8888 | ✓ 1380ms | 否 | ✓ 1965ms | ✓ 1840ms | 否 | http |
| 121.230.8.144:1080 | ✓ 1128ms | ✓ 1604ms | ✓ 1134ms | ✓ 1476ms | ✓ 1392ms | http |
| 61.52.131.172:8443 | ✓ 1129ms | ✓ 1321ms | ✓ 1123ms | 否 | ✓ 1160ms | http |
| 152.32.132.190:7890 | ✓ 1137ms | ✓ 1384ms | ✓ 953ms | ✓ 1125ms | ✓ 925ms | http |

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
