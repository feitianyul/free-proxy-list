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

最后更新：2026-03-28 15:29:34 UTC（2026-03-28 23:29:34 UTC+8）

**代理总数：63**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 63 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 63 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.149.92.147:5001 | ✓ 886ms | 否 | ✓ 871ms | ✓ 1235ms | ✓ 934ms | http |
| 147.161.210.140:8800 | ✓ 1716ms | 否 | ✓ 604ms | ✓ 1271ms | ✓ 1255ms | http |
| 1.231.81.166:3128 | ✓ 1746ms | ✓ 1220ms | ✓ 1179ms | ✓ 1193ms | ✓ 1362ms | http |
| 113.160.132.26:8080 | ✓ 1579ms | ✓ 1270ms | ✓ 932ms | ✓ 1439ms | ✓ 1145ms | http |
| 167.103.115.102:8800 | ✓ 1385ms | 否 | ✓ 848ms | ✓ 1774ms | ✓ 1164ms | http |
| 45.144.28.81:10808 | ✓ 1140ms | 否 | ✓ 1967ms | ✓ 1474ms | ✓ 1608ms | http |
| 167.103.34.108:8800 | ✓ 1409ms | 否 | ✓ 1334ms | ✓ 1816ms | ✓ 1704ms | http |
| 43.99.54.236:5555 | ✓ 986ms | ✓ 906ms | ✓ 579ms | ✓ 735ms | ✓ 595ms | http |
| 103.84.95.54:7890 | ✓ 654ms | 否 | ✓ 740ms | ✓ 940ms | 否 | http |
| 167.103.144.127:8800 | 否 | 否 | ✓ 1332ms | ✓ 1608ms | ✓ 1361ms | http |
| 106.75.15.167:7890 | ✓ 859ms | 否 | 否 | ✓ 1773ms | ✓ 1531ms | http |
| 180.250.219.58:53281 | ✓ 1628ms | 否 | ✓ 1551ms | ✓ 1786ms | ✓ 1707ms | http |
| 168.110.52.228:3128 | 否 | 否 | ✓ 1273ms | ✓ 840ms | ✓ 717ms | http |
| 95.213.217.168:52004 | ✓ 1451ms | 否 | ✓ 1089ms | ✓ 1757ms | ✓ 1343ms | http |
| 167.103.31.122:8800 | ✓ 1401ms | 否 | ✓ 1414ms | 否 | ✓ 1478ms | http |
| 35.225.22.61:80 | ✓ 334ms | ✓ 1331ms | 否 | ✓ 1863ms | ✓ 1206ms | http |
| 186.148.180.46:999 | 否 | 否 | ✓ 1740ms | ✓ 1931ms | ✓ 1665ms | http |
| 20.210.39.153:8561 | ✓ 1299ms | ✓ 932ms | ✓ 483ms | ✓ 788ms | ✓ 626ms | http |
| 20.78.118.91:8561 | ✓ 1298ms | ✓ 830ms | ✓ 577ms | ✓ 797ms | ✓ 621ms | http |
| 20.78.26.206:8561 | ✓ 1296ms | 否 | ✓ 405ms | ✓ 774ms | ✓ 634ms | http |
| 128.199.121.61:9090 | 否 | 否 | ✓ 1953ms | ✓ 1013ms | ✓ 828ms | http |
| 103.113.70.189:1081 | ✓ 433ms | 否 | ✓ 1695ms | ✓ 1249ms | 否 | http |
| 147.161.239.240:8800 | ✓ 762ms | 否 | ✓ 1443ms | ✓ 1732ms | ✓ 1668ms | http |
| 101.43.127.100:8877 | ✓ 881ms | 否 | ✓ 1382ms | 否 | ✓ 1455ms | http |
| 177.234.217.88:999 | ✓ 1676ms | 否 | ✓ 1991ms | 否 | ✓ 1823ms | http |
| 219.117.204.211:7799 | ✓ 496ms | 否 | ✓ 993ms | ✓ 762ms | ✓ 685ms | http |
| 5.104.87.17:8051 | ✓ 1971ms | 否 | ✓ 1466ms | 否 | ✓ 1378ms | http |
| 120.92.212.16:7890 | ✓ 950ms | ✓ 1198ms | ✓ 1324ms | 否 | ✓ 925ms | http |
| 128.199.116.219:9090 | ✓ 1214ms | 否 | ✓ 1416ms | ✓ 1454ms | 否 | http |
| 45.12.151.226:2829 | ✓ 893ms | 否 | ✓ 880ms | 否 | ✓ 1779ms | http |
| 223.16.170.103:80 | ✓ 891ms | 否 | ✓ 901ms | 否 | ✓ 1122ms | http |
| 20.27.11.248:8561 | 否 | 否 | ✓ 989ms | ✓ 911ms | ✓ 604ms | http |
| 16.78.119.130:443 | ✓ 1925ms | ✓ 1386ms | ✓ 1946ms | 否 | 否 | http |
| 208.87.243.199:7878 | 否 | 否 | ✓ 1088ms | ✓ 1048ms | ✓ 907ms | http |
| 103.9.78.2:3128 | ✓ 837ms | 否 | ✓ 904ms | ✓ 1018ms | ✓ 812ms | http |
| 120.92.212.16:8890 | ✓ 1091ms | ✓ 1166ms | 否 | ✓ 1156ms | ✓ 906ms | http |
| 128.199.113.85:9090 | ✓ 1435ms | 否 | ✓ 1425ms | ✓ 1111ms | ✓ 1065ms | http |
| 45.136.131.47:8445 | ✓ 887ms | 否 | ✓ 1179ms | ✓ 944ms | ✓ 1408ms | http |
| 38.145.208.238:8451 | ✓ 523ms | ✓ 1117ms | ✓ 1971ms | ✓ 809ms | ✓ 675ms | http |
| 38.145.208.171:8451 | ✓ 635ms | ✓ 1142ms | ✓ 1636ms | ✓ 786ms | ✓ 1641ms | http |
| 146.190.80.158:9090 | ✓ 1233ms | 否 | ✓ 967ms | ✓ 1373ms | ✓ 1167ms | http |
| 45.136.131.62:8446 | ✓ 1512ms | 否 | ✓ 1561ms | ✓ 961ms | 否 | http |
| 38.145.208.182:8444 | ✓ 1718ms | ✓ 1608ms | 否 | ✓ 1458ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1861ms | 否 | ✓ 1301ms | ✓ 1131ms | ✓ 1417ms | http |
| 222.228.171.92:8080 | ✓ 554ms | ✓ 1868ms | 否 | ✓ 1034ms | ✓ 1397ms | http |
| 85.208.108.43:10808 | ✓ 536ms | 否 | ✓ 1078ms | ✓ 1165ms | ✓ 860ms | http |
| 85.208.108.43:2094 | ✓ 534ms | 否 | ✓ 1079ms | ✓ 1166ms | ✓ 864ms | http |
| 103.82.23.118:5224 | ✓ 1413ms | 否 | ✓ 1317ms | ✓ 1887ms | ✓ 1342ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1317ms | ✓ 1201ms | ✓ 1030ms | http |
| 121.147.253.205:3072 | ✓ 1267ms | ✓ 1595ms | ✓ 1248ms | ✓ 1391ms | ✓ 1355ms | http |
| 45.129.141.143:3128 | ✓ 1297ms | 否 | ✓ 1796ms | 否 | ✓ 1941ms | http |
| 45.136.198.40:3128 | ✓ 1295ms | 否 | ✓ 1783ms | 否 | ✓ 1682ms | http |
| 128.199.114.189:9090 | 否 | 否 | ✓ 1297ms | ✓ 1006ms | ✓ 888ms | http |
| 64.227.76.27:1080 | ✓ 1064ms | 否 | ✓ 1858ms | 否 | ✓ 1949ms | http |
| 137.184.1.87:3128 | ✓ 1056ms | ✓ 855ms | ✓ 1094ms | ✓ 769ms | ✓ 571ms | http |
| 38.145.208.213:8450 | ✓ 615ms | 否 | ✓ 735ms | ✓ 1265ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1729ms | 否 | 否 | ✓ 1822ms | ✓ 1508ms | http |
| 2.56.173.45:10808 | ✓ 1066ms | ✓ 1760ms | ✓ 1067ms | 否 | ✓ 1921ms | http |
| 103.67.46.225:3125 | ✓ 1789ms | 否 | ✓ 1560ms | ✓ 1723ms | ✓ 1474ms | http |
| 61.52.131.172:8443 | ✓ 628ms | ✓ 861ms | ✓ 755ms | ✓ 963ms | ✓ 763ms | http |
| 128.199.254.13:9090 | ✓ 1623ms | 否 | ✓ 1057ms | ✓ 1510ms | ✓ 973ms | http |
| 45.140.147.155:1082 | ✓ 1099ms | ✓ 1389ms | ✓ 1485ms | ✓ 1607ms | ✓ 1278ms | http |
| 160.238.65.8:3128 | ✓ 806ms | 否 | ✓ 1303ms | 否 | ✓ 1766ms | http |

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
