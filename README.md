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

最后更新：2026-04-27 04:22:02 UTC（2026-04-27 12:22:02 UTC+8）

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
| 113.160.132.26:8080 | 否 | ✓ 1662ms | ✓ 1462ms | ✓ 1998ms | ✓ 1363ms | http |
| 217.76.245.80:999 | ✓ 1222ms | ✓ 1755ms | ✓ 1420ms | ✓ 1777ms | ✓ 1686ms | http |
| 139.162.46.62:3128 | ✓ 1133ms | 否 | ✓ 683ms | ✓ 1698ms | ✓ 824ms | http |
| 47.85.51.197:1080 | ✓ 1683ms | 否 | 否 | ✓ 1492ms | ✓ 1997ms | http |
| 80.92.204.47:1081 | ✓ 686ms | 否 | ✓ 1234ms | ✓ 1874ms | ✓ 1241ms | http |
| 218.108.131.186:17890 | ✓ 1292ms | ✓ 1734ms | ✓ 1061ms | 否 | ✓ 1091ms | http |
| 177.93.132.244:3128 | ✓ 1290ms | 否 | ✓ 1833ms | 否 | ✓ 1870ms | http |
| 120.92.212.16:8890 | ✓ 1290ms | ✓ 1579ms | 否 | ✓ 1374ms | 否 | http |
| 103.157.200.126:3128 | ✓ 1608ms | 否 | ✓ 1617ms | ✓ 1886ms | ✓ 1499ms | http |
| 211.95.152.50:45046 | ✓ 1042ms | ✓ 1338ms | ✓ 1149ms | ✓ 1989ms | ✓ 1095ms | http |
| 43.165.195.107:3128 | ✓ 1935ms | ✓ 1407ms | ✓ 1136ms | 否 | 否 | http |
| 34.101.184.164:3128 | ✓ 1873ms | 否 | 否 | ✓ 1708ms | ✓ 1592ms | http |
| 120.92.212.16:7890 | ✓ 992ms | ✓ 1816ms | 否 | ✓ 1701ms | ✓ 1128ms | http |
| 86.104.72.220:1081 | ✓ 832ms | ✓ 1622ms | ✓ 510ms | ✓ 1833ms | ✓ 1526ms | http |
| 59.46.216.131:30001 | ✓ 889ms | 否 | ✓ 1042ms | ✓ 1319ms | ✓ 1152ms | http |
| 210.77.19.6:7890 | ✓ 762ms | ✓ 1026ms | ✓ 771ms | ✓ 1014ms | ✓ 816ms | http |
| 130.61.139.145:3128 | ✓ 1164ms | ✓ 1738ms | 否 | ✓ 1869ms | ✓ 1441ms | http |
| 212.58.132.5:8888 | ✓ 1882ms | 否 | ✓ 1639ms | ✓ 1615ms | ✓ 1250ms | http |
| 117.236.124.166:3128 | ✓ 1940ms | 否 | ✓ 1891ms | 否 | ✓ 1741ms | http |
| 124.121.2.223:8080 | ✓ 1633ms | 否 | 否 | ✓ 1623ms | ✓ 1456ms | http |
| 47.95.231.180:8084 | ✓ 1890ms | 否 | ✓ 1903ms | ✓ 1862ms | ✓ 1190ms | http |
| 101.32.244.83:8080 | ✓ 948ms | 否 | ✓ 926ms | ✓ 1114ms | 否 | http |
| 121.43.196.213:8222 | ✓ 912ms | ✓ 999ms | ✓ 831ms | ✓ 1069ms | ✓ 881ms | http |
| 121.43.196.210:8222 | ✓ 980ms | ✓ 998ms | ✓ 816ms | ✓ 1053ms | ✓ 880ms | http |
| 114.55.226.123:10086 | ✓ 1059ms | ✓ 1424ms | ✓ 1141ms | 否 | 否 | http |
| 8.211.166.184:8081 | ✓ 1365ms | ✓ 1429ms | ✓ 669ms | ✓ 828ms | ✓ 643ms | http |
| 64.188.67.154:1080 | ✓ 1257ms | 否 | ✓ 1531ms | ✓ 1774ms | 否 | http |
| 120.92.108.86:7890 | ✓ 1820ms | 否 | ✓ 1641ms | ✓ 1617ms | ✓ 1857ms | http |
| 47.84.59.16:1080 | ✓ 1264ms | ✓ 1703ms | ✓ 917ms | ✓ 1289ms | ✓ 813ms | http |
| 38.244.52.152:26678 | ✓ 448ms | ✓ 761ms | ✓ 849ms | ✓ 710ms | ✓ 691ms | http |
| 103.172.121.58:7777 | ✓ 1870ms | 否 | 否 | ✓ 1570ms | ✓ 1440ms | http |
| 45.140.147.82:1082 | ✓ 1120ms | ✓ 1564ms | ✓ 1040ms | ✓ 1582ms | ✓ 1530ms | http |
| 45.140.147.82:1081 | ✓ 1113ms | 否 | ✓ 651ms | ✓ 1540ms | 否 | http |
| 86.104.72.219:1081 | ✓ 971ms | ✓ 1164ms | ✓ 1372ms | ✓ 1398ms | ✓ 1003ms | http |
| 47.101.159.19:8899 | ✓ 833ms | ✓ 996ms | ✓ 895ms | ✓ 1037ms | ✓ 902ms | http |
| 62.113.119.14:8080 | ✓ 1404ms | ✓ 1733ms | ✓ 762ms | ✓ 1610ms | ✓ 1245ms | http |
| 1.180.87.146:22300 | ✓ 1684ms | ✓ 1319ms | ✓ 1129ms | ✓ 1294ms | ✓ 1041ms | http |
| 160.250.4.245:1 | ✓ 1571ms | 否 | ✓ 1404ms | ✓ 1466ms | 否 | http |
| 46.101.95.183:8888 | 否 | 否 | ✓ 1682ms | ✓ 1974ms | ✓ 1774ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 887ms | ✓ 1987ms | ✓ 1592ms | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1879ms | ✓ 1566ms | ✓ 785ms | http |
| 8.219.195.129:1080 | ✓ 824ms | ✓ 1722ms | ✓ 837ms | ✓ 982ms | ✓ 816ms | http |
| 45.153.231.229:8080 | ✓ 1255ms | ✓ 1877ms | ✓ 944ms | 否 | ✓ 1483ms | http |
| 38.244.54.190:31168 | ✓ 551ms | ✓ 701ms | ✓ 913ms | ✓ 889ms | ✓ 668ms | http |
| 20.78.118.91:8561 | 否 | 否 | ✓ 484ms | ✓ 871ms | ✓ 913ms | http |
| 20.78.26.206:8561 | 否 | 否 | ✓ 488ms | ✓ 867ms | ✓ 921ms | http |
| 20.210.39.153:8561 | 否 | ✓ 1056ms | ✓ 773ms | ✓ 1129ms | ✓ 1569ms | http |
| 61.52.131.172:8443 | ✓ 882ms | ✓ 1127ms | ✓ 973ms | ✓ 1144ms | ✓ 946ms | http |
| 183.232.248.73:7890 | ✓ 833ms | ✓ 1261ms | 否 | 否 | ✓ 909ms | http |

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
