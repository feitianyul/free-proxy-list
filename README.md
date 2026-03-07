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

最后更新：2026-03-07 11:24:08 UTC（2026-03-07 19:24:08 UTC+8）

**代理总数：68**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 417ms | ✓ 1902ms | ✓ 1168ms | ✓ 1182ms | ✓ 1111ms | http |
| 159.223.42.219:3128 | ✓ 851ms | 否 | ✓ 977ms | ✓ 1136ms | ✓ 911ms | http |
| 46.249.103.192:443 | ✓ 1122ms | 否 | ✓ 1462ms | ✓ 1673ms | ✓ 1981ms | http |
| 91.107.175.112:10801 | 否 | 否 | ✓ 1867ms | ✓ 1257ms | ✓ 1509ms | http |
| 61.72.221.234:3128 | 否 | 否 | ✓ 1486ms | ✓ 1567ms | ✓ 1697ms | http |
| 217.76.245.80:999 | ✓ 531ms | ✓ 1493ms | ✓ 1154ms | ✓ 1704ms | ✓ 1300ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 316ms | ✓ 1197ms | ✓ 1091ms | http |
| 14.56.107.244:3128 | ✓ 1221ms | ✓ 1823ms | ✓ 898ms | 否 | ✓ 867ms | http |
| 125.128.12.144:3128 | ✓ 1220ms | ✓ 1487ms | ✓ 724ms | ✓ 1290ms | 否 | http |
| 61.72.110.94:3128 | 否 | ✓ 1571ms | ✓ 1621ms | ✓ 1727ms | ✓ 1947ms | http |
| 103.139.138.194:3128 | ✓ 1270ms | 否 | ✓ 1754ms | ✓ 1543ms | 否 | http |
| 203.205.49.2:10239 | 否 | 否 | ✓ 1637ms | ✓ 1385ms | ✓ 1359ms | http |
| 120.92.212.16:8890 | ✓ 1062ms | 否 | ✓ 1514ms | ✓ 1285ms | ✓ 1057ms | http |
| 46.183.25.8:443 | 否 | 否 | ✓ 350ms | ✓ 1073ms | ✓ 825ms | http |
| 1.231.81.166:3128 | ✓ 886ms | ✓ 1844ms | ✓ 880ms | ✓ 1102ms | ✓ 760ms | http |
| 34.101.184.164:3128 | ✓ 1783ms | 否 | ✓ 1755ms | ✓ 1371ms | ✓ 1447ms | http |
| 121.128.121.54:3128 | ✓ 744ms | 否 | ✓ 1558ms | ✓ 1048ms | ✓ 1936ms | http |
| 193.168.173.136:443 | 否 | 否 | ✓ 1006ms | ✓ 1792ms | ✓ 1626ms | http |
| 168.235.110.63:3128 | ✓ 665ms | 否 | ✓ 1709ms | ✓ 1152ms | ✓ 1871ms | http |
| 162.248.165.72:1080 | ✓ 536ms | 否 | ✓ 1762ms | ✓ 1792ms | ✓ 1747ms | http |
| 91.193.240.157:9877 | ✓ 1082ms | 否 | ✓ 1101ms | 否 | ✓ 1873ms | http |
| 81.70.169.194:80 | ✓ 1156ms | 否 | ✓ 1393ms | 否 | ✓ 1116ms | http |
| 101.43.255.96:80 | 否 | 否 | ✓ 1326ms | ✓ 1402ms | ✓ 1072ms | http |
| 115.231.181.40:8128 | ✓ 1023ms | ✓ 1418ms | ✓ 1245ms | 否 | 否 | http |
| 185.115.74.185:8080 | ✓ 923ms | ✓ 1879ms | ✓ 1443ms | 否 | 否 | http |
| 62.113.119.14:8080 | ✓ 1076ms | ✓ 1609ms | ✓ 1071ms | ✓ 1520ms | ✓ 1192ms | http |
| 103.84.95.54:7890 | ✓ 1204ms | 否 | ✓ 1805ms | ✓ 1245ms | ✓ 1654ms | http |
| 120.92.212.16:7890 | ✓ 1835ms | ✓ 1572ms | ✓ 1057ms | ✓ 1341ms | ✓ 1308ms | http |
| 190.9.109.205:999 | ✓ 682ms | ✓ 1461ms | ✓ 1180ms | ✓ 1383ms | ✓ 1225ms | http |
| 190.9.109.199:999 | ✓ 789ms | ✓ 1535ms | ✓ 1292ms | ✓ 1275ms | 否 | http |
| 116.80.82.232:3172 | ✓ 1997ms | 否 | 否 | ✓ 1943ms | ✓ 1890ms | http |
| 178.236.245.17:3128 | ✓ 629ms | ✓ 1747ms | ✓ 1429ms | ✓ 1660ms | ✓ 1254ms | http |
| 178.236.245.59:3128 | ✓ 600ms | 否 | ✓ 1193ms | ✓ 1698ms | ✓ 1286ms | http |
| 167.172.69.123:8080 | 否 | 否 | ✓ 1910ms | ✓ 1611ms | ✓ 964ms | http |
| 167.172.69.123:80 | ✓ 1626ms | 否 | ✓ 1493ms | ✓ 1305ms | ✓ 1114ms | http |
| 14.225.217.30:7890 | ✓ 837ms | ✓ 1964ms | 否 | 否 | ✓ 1133ms | http |
| 103.86.131.62:80 | ✓ 1663ms | 否 | 否 | ✓ 1873ms | ✓ 1238ms | http |
| 45.140.147.155:1081 | ✓ 456ms | 否 | ✓ 841ms | ✓ 1431ms | 否 | http |
| 85.9.195.140:1080 | 否 | 否 | ✓ 925ms | ✓ 1854ms | ✓ 1225ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1861ms | 否 | ✓ 1139ms | ✓ 856ms | http |
| 175.0.74.111:10808 | 否 | 否 | ✓ 1091ms | ✓ 1399ms | ✓ 1134ms | http |
| 103.215.36.88:18332 | ✓ 1080ms | 否 | ✓ 1252ms | 否 | ✓ 1149ms | http |
| 210.223.44.230:3128 | ✓ 1608ms | ✓ 923ms | ✓ 1047ms | 否 | ✓ 773ms | http |
| 101.32.244.83:8080 | ✓ 1775ms | 否 | ✓ 1036ms | ✓ 1364ms | ✓ 1334ms | http |
| 121.43.196.213:8222 | ✓ 1033ms | ✓ 1181ms | ✓ 1001ms | ✓ 1257ms | ✓ 955ms | http |
| 121.43.196.210:8222 | ✓ 1126ms | ✓ 1182ms | ✓ 914ms | ✓ 1224ms | ✓ 971ms | http |
| 114.55.226.123:10086 | ✓ 1374ms | 否 | ✓ 1119ms | ✓ 1345ms | ✓ 1080ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1349ms | ✓ 1273ms | ✓ 1375ms | http |
| 192.166.82.55:1080 | ✓ 1044ms | ✓ 1774ms | ✓ 907ms | ✓ 1267ms | 否 | http |
| 38.47.97.22:6005 | ✓ 1665ms | 否 | 否 | ✓ 1677ms | ✓ 1498ms | http |
| 202.129.206.239:3128 | ✓ 1508ms | 否 | ✓ 1742ms | 否 | ✓ 1993ms | http |
| 45.186.6.104:3128 | ✓ 1299ms | ✓ 1962ms | ✓ 1897ms | 否 | 否 | http |
| 37.187.109.70:10111 | ✓ 1731ms | ✓ 1808ms | ✓ 1989ms | 否 | 否 | http |
| 116.80.82.224:3172 | ✓ 1673ms | 否 | 否 | ✓ 1987ms | ✓ 1885ms | http |
| 8.137.112.117:3128 | ✓ 1528ms | 否 | 否 | ✓ 1599ms | ✓ 1212ms | http |
| 106.14.203.63:3333 | ✓ 904ms | ✓ 1145ms | ✓ 912ms | ✓ 1248ms | ✓ 1082ms | http |
| 66.228.47.125:110 | ✓ 695ms | 否 | ✓ 1041ms | ✓ 1271ms | ✓ 891ms | http |
| 45.136.198.40:3128 | ✓ 775ms | ✓ 1749ms | ✓ 1992ms | 否 | ✓ 1692ms | http |
| 194.59.204.87:9080 | ✓ 1171ms | ✓ 1628ms | ✓ 1746ms | 否 | ✓ 1754ms | http |
| 103.39.51.190:8080 | ✓ 1821ms | 否 | ✓ 1329ms | ✓ 1815ms | ✓ 1564ms | http |
| 165.227.5.10:8888 | ✓ 192ms | 否 | ✓ 669ms | ✓ 1002ms | ✓ 637ms | http |
| 125.128.12.14:3128 | 否 | 否 | ✓ 1128ms | ✓ 1087ms | ✓ 833ms | http |
| 172.212.68.37:3128 | ✓ 406ms | ✓ 1563ms | ✓ 572ms | ✓ 1580ms | ✓ 1281ms | http |
| 1.234.153.14:80 | ✓ 1470ms | ✓ 1376ms | ✓ 1422ms | ✓ 1962ms | 否 | http |
| 120.55.163.237:10086 | ✓ 1708ms | ✓ 1170ms | ✓ 1004ms | ✓ 1203ms | ✓ 1000ms | http |
| 121.126.185.63:25152 | ✓ 1593ms | 否 | ✓ 1354ms | 否 | ✓ 1863ms | http |
| 88.80.150.82:8080 | ✓ 837ms | 否 | ✓ 1047ms | 否 | ✓ 1369ms | https |
| 103.215.36.88:19217 | ✓ 1406ms | ✓ 1910ms | ✓ 1106ms | ✓ 1611ms | ✓ 1408ms | http |

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
