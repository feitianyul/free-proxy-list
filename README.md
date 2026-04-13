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

最后更新：2026-04-13 16:07:13 UTC（2026-04-14 00:07:13 UTC+8）

**代理总数：49**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 49 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 49 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1525ms | ✓ 1276ms | ✓ 804ms | ✓ 1117ms | ✓ 1034ms | http |
| 167.103.115.102:8800 | ✓ 1759ms | 否 | ✓ 1157ms | 否 | ✓ 1127ms | http |
| 113.160.132.26:8080 | ✓ 1023ms | 否 | ✓ 1330ms | ✓ 1357ms | 否 | http |
| 167.103.34.108:8800 | ✓ 1252ms | 否 | ✓ 1264ms | ✓ 1500ms | 否 | http |
| 167.103.144.127:8800 | ✓ 1064ms | 否 | ✓ 1234ms | ✓ 1822ms | 否 | http |
| 159.223.225.118:8888 | 否 | 否 | ✓ 701ms | ✓ 1632ms | ✓ 1396ms | http |
| 45.167.125.21:999 | ✓ 686ms | ✓ 1952ms | 否 | ✓ 1695ms | ✓ 1473ms | http |
| 212.58.132.5:8888 | ✓ 1141ms | 否 | ✓ 1473ms | ✓ 1561ms | ✓ 1218ms | http |
| 167.103.31.122:8800 | ✓ 1498ms | 否 | ✓ 1530ms | ✓ 1796ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1093ms | 否 | 否 | ✓ 1832ms | ✓ 1262ms | http |
| 36.103.198.235:7890 | ✓ 1220ms | ✓ 1990ms | ✓ 1282ms | 否 | ✓ 1347ms | http |
| 218.108.131.186:17890 | ✓ 1240ms | ✓ 1263ms | ✓ 984ms | ✓ 1256ms | 否 | http |
| 147.161.239.240:8800 | ✓ 714ms | ✓ 1619ms | ✓ 1187ms | 否 | ✓ 1503ms | http |
| 43.156.132.113:3128 | ✓ 1719ms | 否 | ✓ 1232ms | 否 | ✓ 1643ms | http |
| 210.223.44.230:3128 | ✓ 1750ms | ✓ 1499ms | 否 | ✓ 1416ms | ✓ 848ms | http |
| 116.203.112.97:3128 | ✓ 1722ms | ✓ 1955ms | ✓ 820ms | ✓ 1322ms | ✓ 1225ms | http |
| 192.71.213.85:9090 | ✓ 1422ms | 否 | ✓ 1673ms | ✓ 1979ms | 否 | http |
| 79.132.136.58:3128 | ✓ 482ms | 否 | ✓ 474ms | ✓ 1566ms | ✓ 1315ms | http |
| 103.3.246.71:3128 | 否 | 否 | ✓ 1271ms | ✓ 1757ms | ✓ 1834ms | http |
| 103.39.51.207:8080 | ✓ 1533ms | 否 | 否 | ✓ 1911ms | ✓ 1930ms | http |
| 35.225.22.61:80 | ✓ 780ms | 否 | ✓ 983ms | ✓ 1263ms | ✓ 879ms | http |
| 168.110.52.228:3128 | 否 | 否 | ✓ 1586ms | ✓ 954ms | ✓ 1872ms | http |
| 185.76.240.64:10001 | ✓ 1259ms | 否 | ✓ 1819ms | 否 | ✓ 1615ms | http |
| 171.227.167.109:1004 | ✓ 1009ms | 否 | 否 | ✓ 1430ms | ✓ 1084ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1172ms | ✓ 1412ms | ✓ 1416ms | http |
| 45.140.147.155:1082 | ✓ 484ms | ✓ 1281ms | ✓ 1540ms | ✓ 1517ms | ✓ 1226ms | http |
| 5.104.87.17:8051 | ✓ 1442ms | 否 | ✓ 1661ms | ✓ 1334ms | ✓ 1020ms | http |
| 120.92.108.86:7890 | ✓ 1793ms | 否 | 否 | ✓ 1850ms | ✓ 1517ms | http |
| 34.71.229.255:3128 | ✓ 980ms | 否 | ✓ 1107ms | ✓ 1419ms | ✓ 1376ms | http |
| 121.130.199.80:3128 | 否 | 否 | ✓ 1496ms | ✓ 1473ms | ✓ 1635ms | http |
| 137.59.47.73:3128 | 否 | 否 | ✓ 1477ms | ✓ 1444ms | ✓ 1491ms | http |
| 8.219.195.129:1080 | ✓ 953ms | ✓ 1923ms | ✓ 874ms | ✓ 1175ms | ✓ 947ms | http |
| 107.172.102.234:40621 | ✓ 1100ms | ✓ 1844ms | 否 | ✓ 1722ms | 否 | http |
| 162.240.154.26:3128 | ✓ 1967ms | ✓ 1342ms | ✓ 1165ms | 否 | ✓ 985ms | http |
| 205.164.46.6:3082 | ✓ 1870ms | 否 | 否 | ✓ 1910ms | ✓ 1365ms | http |
| 194.67.99.223:1080 | ✓ 618ms | ✓ 1867ms | ✓ 778ms | ✓ 1805ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1020ms | 否 | ✓ 907ms | ✓ 1297ms | ✓ 1017ms | http |
| 46.30.46.133:3128 | ✓ 1083ms | 否 | ✓ 1578ms | ✓ 1381ms | ✓ 1017ms | http |
| 157.245.194.13:8888 | ✓ 1036ms | 否 | ✓ 1398ms | ✓ 1994ms | ✓ 1866ms | http |
| 185.76.240.61:10001 | ✓ 1076ms | 否 | ✓ 1333ms | 否 | ✓ 1981ms | http |
| 61.52.131.172:8443 | ✓ 1004ms | ✓ 1321ms | ✓ 1135ms | ✓ 1420ms | ✓ 1091ms | http |
| 130.61.30.221:8080 | 否 | 否 | ✓ 1676ms | ✓ 1975ms | ✓ 1516ms | http |
| 101.43.127.100:8877 | ✓ 1078ms | 否 | ✓ 1000ms | ✓ 1926ms | ✓ 969ms | http |
| 62.113.119.14:8080 | ✓ 988ms | 否 | ✓ 1313ms | 否 | ✓ 1975ms | http |
| 185.191.236.162:3128 | ✓ 948ms | ✓ 1763ms | ✓ 1461ms | 否 | 否 | http |
| 36.141.21.200:7890 | ✓ 1032ms | ✓ 1940ms | ✓ 1111ms | ✓ 1365ms | 否 | http |
| 20.210.76.104:8561 | ✓ 644ms | ✓ 1193ms | ✓ 708ms | ✓ 995ms | ✓ 720ms | http |
| 20.210.76.178:8561 | ✓ 636ms | ✓ 1180ms | ✓ 751ms | ✓ 944ms | ✓ 738ms | http |
| 47.252.41.213:443 | ✓ 282ms | ✓ 1193ms | ✓ 347ms | ✓ 1122ms | ✓ 1035ms | http |

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
